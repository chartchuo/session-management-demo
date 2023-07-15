import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'api.dart';
import 'home_page.dart';
import 'login_page.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  var api = Api();
  await api.session.load();
  runApp(const MyApp());
}

final GoRouter _router = GoRouter(
  initialLocation: Api().isLogin ? '/' : 'login',
  routes: <RouteBase>[
    GoRoute(
      path: '/',
      builder: (BuildContext context, GoRouterState state) {
        return const MyHomePage();
      },
      routes: <RouteBase>[
        GoRoute(
          path: 'login',
          builder: (BuildContext context, GoRouterState state) {
            return const LoginPage();
          },
        ),
      ],
    ),
  ],
);

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      routerConfig: _router,
    );
  }
}
