import 'package:flutter/material.dart';

import 'api.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  String msg = '';

  void _incrementCounter() async {
    var api = Api();
    if (!api.isLogin) {
      await api.login('admin', 'admin');
    }
    var data = await api.userHello();
    setState(() {
      msg = data;
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
            const Text(
              'You have pushed the button this many times:',
            ),
            Text(
              msg,
              // style: Theme.of(context).textTheme.headlineMedium,
            ),
            Text('next refresh time: ${Api().session.atExpDate}'),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
