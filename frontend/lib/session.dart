import 'dart:convert';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:frontend/model.dart';
import 'package:jwt_decoder/jwt_decoder.dart';

class Session {
  final storage = const FlutterSecureStorage();
  Tokens tokens = Tokens('', '');
  DateTime rtExpDate = DateTime.now();
  DateTime atExpDate = DateTime.now();
  bool isLogin = false;

  bool rtExp() => rtExpDate.isBefore(DateTime.now());
  bool atExp() => atExpDate.isBefore(DateTime.now());

  load() async {
    var str = await storage.read(key: 'tokens');
    if (str != null) {
      tokens = Tokens.fromJson(jsonDecode(str));
    }
    _updateExpire();
  }

  save() async {
    var str = jsonEncode(tokens.toJson());
    await storage.write(key: 'tokens', value: str);
    _updateExpire();
  }

  void refreshTokenRotate(Tokens newTokens) {
    tokens = newTokens;
    save();
  }

  void _updateExpire() {
    if (tokens.refreshToken == '') {
      clearSession();
      return;
    }

    rtExpDate = JwtDecoder.getExpirationDate(tokens.refreshToken);
    atExpDate = JwtDecoder.getExpirationDate(tokens.accessToken)
        .subtract(const Duration(seconds: 10)); // sub 10 secs for safety
    isLogin = true;

    if (rtExp()) {
      clearSession();
    }
  }

  void clearSession() {
    isLogin = false;
    tokens = tokens = Tokens('', '');
  }
}
