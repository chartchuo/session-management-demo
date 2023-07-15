import 'package:json_annotation/json_annotation.dart';

part 'model.g.dart';

// cmd
// dart run build_runner build --delete-conflicting-outputs

@JsonSerializable()
class Tokens {
  @JsonKey(name: 'refresh_token')
  final String refreshToken;
  @JsonKey(name: 'access_token')
  final String accessToken;

  Tokens(this.refreshToken, this.accessToken);

  factory Tokens.fromJson(Map<String, dynamic> json) => _$TokensFromJson(json);
  Map<String, dynamic> toJson() => _$TokensToJson(this);
}
