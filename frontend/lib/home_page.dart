import 'dart:async';

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
  String userHelloMsg = '';
  var api = Api();

  void _userHello() async {
    var data = await api.userHello();
    setState(() {
      userHelloMsg = data;
    });
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
              onPressed: _userHello,
              child: const Text('Get user hello'),
            ),
            Text(
              userHelloMsg,
              // style: Theme.of(context).textTheme.headlineMedium,
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
}
