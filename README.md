<div align="center">
  <a href="https://gabrielbarker.github.io/identico/">
    <img src="https://identico.herokuapp.com/?quadrant=8" width="200" height="200">
  </a>

  <h1>identico</h1>

  <a href="https://travis-ci.com/gabrielbarker/identico">
    <img src="https://travis-ci.com/gabrielbarker/identico.svg?branch=main"/>
  </a>
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/language-go-00ADD8"/>
  </a>
  <a href="https://github.com/gabrielbarker/identico/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/gabrielbarker/clyp" />
  </a>

<i>A RESTful API for identicon SVGs</i>

</div>


Identico produces unique two-colour identicon SVG's based on an input message! 

You can use Identico to create fun graphics for encrypted words or to create a personalised profile image using your username, ensuring you have a graphic unique to you! 

## Features!

-	Generate a new graphic by inserting a new word into the _name_ box. 

-	Save your theme with a simple right click, then _save as SVG image_. 

-	Watch as your design changes in real time according to your input word via the Identico web app.

-	Change the complexity of the graphic by clicking the _complexity_ drop down and navigating to a different number. 

## Using The API!
The Identico API can be accessed using the URL
```
https://identico.herokuapp.com/
```
### Parameters 
##### Quadrant
The quadrant parameter dictates how many pixels each quarter of the graphic is before it is mirrored to create the full square graphic. Ie. The width is 2X the quadrant. In the case no quadrant value is given by the user the value 8 is used as a default. 

__Example of URL using only the quadrant parameter where the value is 10__
```
https://identico.herokuapp.com/?quadrant=10
```

##### Message
The message parameter is the line of text that is converted to create the identicon image. It is encrypted using a hash function and is then converted into a graphic. In the case no message is input by the user the current datetime is used as a default. 

__Example of URL using only a message where the message is "Hello World"__
```
https://identico.herokuapp.com/?message=Hello%20World
```

## Using The Web App!
- The name box refers to the message parameter. Simply enter a message to see its real time graphic conversion.
- The complexity box refers to the quadrant parameter. The values range from 4 through to 64, where 64 results in a pattern created from more pixels! 
- You can copy your unique URL from the URL box to link directly to your specified message and complexity pattern. 

__Example of standard URL using a quadrant and a message where the message is "Hello World" and the quadrant value is 10__
```
https://identico.herokuapp.com/?message=hello%20world&quadrant=10
```

## Misc!
If you find that there are issues or improvements for Identico then please feel free to [file an issue!](https://github.com/gabrielbarker/identico/issues)

__This application is a fun way to generate a new graphic for a profile when you end up with an automated one that just sucks! It is intended as a solution to forever being stuck between the choice of your actual face or a pre-picked design that just doesn’t suit your personality.__

## License!
[MIT](./LICENSE)
