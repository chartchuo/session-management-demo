import 'dart:io';

import 'package:dio/dio.dart';
import 'package:frontend/model.dart';
import 'package:frontend/session.dart';

class Api {
  late Dio dio;
  var session = Session();
  var baseOptions = BaseOptions(
    baseUrl: 'http://10.0.2.2:8000',
    connectTimeout: const Duration(seconds: 60),
    receiveTimeout: const Duration(seconds: 10),
  );
  bool get isLogin => session.isLogin;
  set isLogin(bool b) => session.isLogin = b;

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
    session.refreshTokenRotate(Tokens.fromJson(response.data));
  }

  tokenInterceptor(
      RequestOptions options, RequestInterceptorHandler handler) async {
    // not login
    if (!isLogin) {
      handler.next(options);
      return;
    }

    // session timeout
    if (session.rtExp()) {
      handler.reject(DioException(requestOptions: options));
      isLogin = false;
      return;
    }

    // access timeout
    if (session.atExp()) {
      //special dio instance for refresh new token
      var response = await Dio(baseOptions).get(
        '/refresh_token',
        options: Options(
          headers: {
            'Authorization': 'Bearer ${session.tokens.refreshToken}',
          },
        ),
      );
      if (response.statusCode != HttpStatus.ok) {
        handler.reject(DioException(requestOptions: options));
        return;
      }
      session.tokens = Tokens.fromJson(response.data);
      session.save();
    }

    //normal case
    options.headers.addAll({
      'Authorization': 'Bearer ${session.tokens.accessToken}',
    });

    handler.next(options);
  }

  logout() {
    session.clearSession();
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
