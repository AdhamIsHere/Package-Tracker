import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-splash-screen',
  templateUrl: './splash-screen.component.html',
  styleUrls: ['./splash-screen.component.css']
})
export class SplashScreenComponent implements OnInit {
  displayedText = '';
  fullText = 'Package Tracker';
  charIndex = 0;

  ngOnInit(): void {
    this.animateText();
  }

  animateText() {
    const interval = setInterval(() => {
      if (this.charIndex < this.fullText.length) {
        this.displayedText += this.fullText[this.charIndex];
        this.charIndex++;
      } else {
        clearInterval(interval);
      }
    }, 100);  // Adjust speed of character appearance (100 ms between characters)
  }
}
