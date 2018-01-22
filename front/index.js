token = 'eyJhbGciOiJSUzI1NiJ9.eyJyb2xlcyI6WyJST0xFX1VTRVIiXSwidXNlcm5hbWUiOiJ0ZXN0MiIsImlhdCI6MTUxNjY1NzkzOCwiZXhwIjoxNTE2NjYxNTM4fQ.GyOOdRnEyhzEhcEjM7Pka58XMUaW2SZXUo8ETVpbxHZdSlyxn76DjL407_cHw9Auwe2JOqK4f-tIaU-J-uS-Rgz5AAdyJ9vf2zavNTUhBqJzhHA6wJlxMazsfsZhDl3nrzPxZmAwUqTh0oBpBBfDJhLpj062WM130KHtGxZ51wetrbBoESXzwTkelTq6WGCsg-J8chPTOHSB1-Ob5fPtuIZTfmRHYAKK9cYUlE3z3mxpPglHyQV9VIep71qLztIGwXmmBnmAwYqhBFq8mAUPSnD8HBMUzcXQGnXeCGyaWOea5h57XnHJ6RwyncURqF2f7Gkba4js9jqfX64fA-pf-Umz1Qk_k5X7-WqAcoGGGJhkAAgUhZLMbg80ufV82L_KB3OFptiKzuD0ZeMJQ5VYpXuNoD4gFGf9syVKBeS1Abz4GjvNCqp73MSalll3D5YBJdHhy7xsHAYWYUE_CSOj2TCX2fBvnTVqAaqqc9q4HD-Tq97qi1S4zaGP1vyy9PbAxkwaO6ECv1sQLSZjUqKSE2ZdnVvoZz_9UAQkavd_hFKLkHx2t_-99BH5zW9ry_Mpj9ZAkncArlFyquU7H7NiQDIFye71PUjUXFLHuwO_N_MRWZuHuRVzDlBGElH6N68gMbpBUTh9EV4Ugq-953Ioqamo_c2p-ZIiTKKHI7ntTIE';
url = 'ws://localhost:4000/websocket?token='+token;

ws = new WebSocket(url);
ws.onopen = function (event) {
   console.log("Here's some text that the server is urgently awaiting!"); 
};

var details = {
    'userName': 'test@gmail.com',
    'password': 'Password!',
    'grant_type': 'password'
};

var formBody = [];
for (var property in details) {
  var encodedKey = encodeURIComponent(property);
  var encodedValue = encodeURIComponent(details[property]);
  formBody.push(encodedKey + "=" + encodedValue);
}
formBody = formBody.join("&");

fetch('https://example.com/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
  },
  body: formBody
})

fetch('products.json')
.then(function(response) { return response.json(); })
.then(function(json) {
  for(var i = 0; i < json.products.length; i++) {
    var listItem = document.createElement('li');
    listItem.innerHTML = '<strong>' + json.products[i].Name + '</strong>';
    listItem.innerHTML +=' can be found in ' + json.products[i].Location + '.';
    listItem.innerHTML +=' Cost: <strong>£' + json.products[i].Price + '</strong>';
    myList.appendChild(listItem);
  }
});