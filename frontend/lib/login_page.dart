import 'package:flutter/material.dart';
import 'package:frontend/api.dart';
import 'package:go_router/go_router.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  var usernameCtrl = TextEditingController();
  var passwordCtrl = TextEditingController();

  var api = Api();

  void _login() async {
    showLoaderDialog(context);
    var ok = await api.login(usernameCtrl.text, passwordCtrl.text);
    Navigator.pop(context);
    if (ok) {
      context.go('/');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text('try username/password test/test or admin/admin'),
          TextField(controller: usernameCtrl),
          TextField(controller: passwordCtrl),
          ElevatedButton(onPressed: _login, child: const Text('Sign on')),
        ],
      ),
    );
  }
}

showLoaderDialog(BuildContext context) {
  AlertDialog alert = AlertDialog(
    content: Row(
      children: [
        const CircularProgressIndicator(),
        Container(
            margin: const EdgeInsets.only(left: 7),
            child: const Text("Loading...")),
      ],
    ),
  );
  showDialog(
    barrierDismissible: false,
    context: context,
    builder: (BuildContext context) {
      return alert;
    },
  );
}
