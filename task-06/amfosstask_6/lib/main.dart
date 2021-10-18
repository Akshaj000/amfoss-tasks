import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:introduction_screen/introduction_screen.dart';


void main() => runApp(MaterialApp(
  title: "App",
  home: Myapp(),
));


class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('WELCOME'),
        backgroundColor: Colors.brown[100],
        elevation: 0,
        ),
      body: Padding(
        padding: const EdgeInsets.fromLTRB(40.0,20.0,40.0,0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Image.asset('assets/welcome.png'),
            Text(
              'AKSHAJ.S.R',
            style: TextStyle(
              fontSize: 30.0,
              fontWeight: FontWeight.bold,
            )),
            SizedBox(height: 10.0),
            Center(
              child: Text('          Im a first year AIE student studying at Amrita,Amritapuri. I dont have a lot of hobbies but as a get away i enjoy traveling if possible and take pictures. Ive always been interested in computer science coz i dont have to get out of my comfortzone a lot but idk.... hope its gonna change. I chose ai since its a growing field and i hope i could be a building block of it. My main intention is to gain enough skill set for my future jobs or maybe studies. I wish i could bring inovations and thing a little outside the box',
              style: TextStyle(
                fontSize: 18.0,
              )),
            ),
          ],
        ),
      ),
      );
  }
}

class Myapp extends StatelessWidget {

  List<PageViewModel> getpages(){
    return[
      PageViewModel(
        image: Image.asset('assets/page1.png'),
        title: 'YOGA SURGE',
        body: 'Welcome to yoga world',
        footer: Text('Inhale the future, exhale the past'),
      ),
      PageViewModel(
        image: Image.asset('assets/page2.png'),
        title: 'Healthy Freaks Exercise',
        body: 'Letting go is hardest asana',
      ),
      PageViewModel(
        image: Image.asset('assets/page3.png'),
        title: 'Cycling',
        body: 'You cannot always control what \n goes on outside. But you can \n always control what goes on inside',
      ),
      PageViewModel(
        image: Image.asset('assets/page4.png'),
        title: 'Meditation',
        body: 'The longest journey of any person \n is the journey inward.',
      ),
    ];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: IntroductionScreen(
        globalBackgroundColor: Colors.brown[100],
        showSkipButton: true,
        skip: Text('Skip'),
        showNextButton: true,
        next: Text(''),
        nextColor: Colors.orange[50],
        skipColor: Colors.black,
        done: Text('Get Started', style: TextStyle(color: Colors.black)),
        onDone: () {
          Navigator.push(context,MaterialPageRoute(builder: (context) => Home()));
          },
        pages: getpages(),
        dotsDecorator: DotsDecorator(
          size: const Size.square(9.0),
          activeSize: const Size(18.0, 9.0),
          activeShape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(5.0)),
          color: Colors.grey,
          activeColor: Colors.black,
        ),
        ),
      );
  }
}


