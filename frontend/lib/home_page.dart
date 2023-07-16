import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'api.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var now = StreamController<DateTime>();
  var api = Api();

  void _get(String path) async {
    var data = await api.get(path);
    var text = const JsonEncoder.withIndent('  ').convert(data);
    await _showDialog(text);
  }

  @override
  void initState() {
    super.initState();
    Timer.periodic(const Duration(seconds: 1), (timer) {
      now.add(DateTime.now());
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('Home page'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text('next refresh time: ${Api().session.atExpDate}'),
            StreamBuilder(
                stream: now.stream,
                builder: (c, sh) {
                  if (sh.hasData) {
                    return Text(
                        'time to refresh: ${sh.data?.difference(api.session.atExpDate).inSeconds}');
                  }
                  return const Text('');
                }),
            ElevatedButton(
              onPressed: () => _get('/user/test/hello'),
              child: const Text('Get /user/test/hello'),
            ),
            ElevatedButton(
              onPressed: () => _get('/user/abc/hello'),
              child: const Text('Get /user/abc/hello'),
            ),
            ElevatedButton(
              onPressed: () => _get('/admin/hello'),
              child: const Text('Get /admin/hello'),
            ),
            ElevatedButton(
                onPressed: () {
                  Api().logout();
                  context.go('/login');
                },
                child: const Text('Sign off'))
          ],
        ),
      ),
    );
  }

  Future<void> _showDialog(String content) async {
    return showDialog<void>(
      context: context,
      barrierDismissible: false, // user must tap button!
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Response'),
          content: SingleChildScrollView(
            child: ListBody(
              children: <Widget>[
                Text(content),
              ],
            ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('Ok'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }
}
