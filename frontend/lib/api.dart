import 'dart:io';

import 'package:dio/dio.dart';
import 'package:frontend/model.dart';
import 'package:jwt_decoder/jwt_decoder.dart';

class Api {
  late Dio dio;
  var baseOptions = BaseOptions(
    baseUrl: 'http://10.0.2.2:8000',
    connectTimeout: const Duration(seconds: 60),
    receiveTimeout: const Duration(seconds: 10),
  );
  Tokens tokens = Tokens('', '');
  DateTime sessionExp = DateTime.now();
  DateTime refreshTime = DateTime.now();
  bool isLogin = false;

  // singleton pattern
  static final Api _api = Api._constructor();
  factory Api() {
    return _api;
  }

  Api._constructor() {
    dio = Dio(baseOptions);
    dio.interceptors.add(InterceptorsWrapper(
      onRequest: tokenInterceptor,
    ));
    tokens = Tokens('', '');
  }

  Future login(String username, String password) async {
    var response = await dio.post(
      '/login',
      data: {
        'username': username,
        'password': password,
      },
    );
    if (response.statusCode != HttpStatus.ok) {
      throw response.data;
    }
    tokens = Tokens.fromJson(response.data);
    _updateExpire();
    isLogin = true;
  }

  tokenInterceptor(
      RequestOptions options, RequestInterceptorHandler handler) async {
    // not login
    if (!isLogin) {
      handler.next(options);
      return;
    }
    var now = DateTime.now();
    // session timeout
    if (sessionExp.isBefore(now)) {
      handler.reject(DioException(requestOptions: options));
      isLogin = false;
      return;
    }

    // access timeout
    if (refreshTime.isBefore(now)) {
      //special dio instance for refresh new token
      var response = await Dio(baseOptions).get(
        '/refresh_token',
        options: Options(
          headers: {
            'Authorization': 'Bearer ${tokens.refreshToken}',
          },
        ),
      );
      if (response.statusCode != HttpStatus.ok) {
        handler.reject(DioException(requestOptions: options));
        return;
      }
      tokens = Tokens.fromJson(response.data);
      _updateExpire();
    }

    //normal case
    options.headers.addAll({
      'Authorization': 'Bearer ${tokens.accessToken}',
    });

    handler.next(options);
  }

  logout() {
    isLogin = false;
    tokens = tokens = Tokens('', '');
  }

  void _updateExpire() {
    sessionExp = JwtDecoder.getExpirationDate(tokens.refreshToken);
    refreshTime = JwtDecoder.getExpirationDate(tokens.accessToken)
        .subtract(const Duration(seconds: 10));
  }

  Future<String> userHello() async {
    var response = await dio.get(
      '/user/test/hello',
    );
    if (response.statusCode != HttpStatus.ok) {
      throw response.data;
    }
    return Future.value(response.data.toString());
  }
}
