document.getElementById('signin-button').onclick = function(){
  var input = {
    'username': document.getElementById('inputUsername').value,
    'password': document.getElementById('inputPassword').value,
  };

  var div = document.getElementById('user-info');
  div.innerHTML += ' '+input.username;

  var formBody = [];
  for (var property in input) {
  var encodedKey = encodeURIComponent(property);
  var encodedValue = encodeURIComponent(input[property]);
  formBody.push(encodedKey + "=" + encodedValue);
  }
  formBody = formBody.join("&");

  fetch('http://symfony.dev:8007/login_check', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    },
    body: formBody
  }).then(function(response) {
  return response.json();
  }).then(function(json) {
    token = json.token;

    fetch('http://symfony.dev:8007/conversations/1', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json;charset=UTF-8',
        'Authorization': 'Bearer '+token
      }
    }).then(function(response) {
      return response.json();
    }).then(function(json) {
      var messages = json.messages;
      messages.reverse();
      messages.forEach(message => {
        var div = document.createElement('div');
        div.id = 'block';
        div.className = 'block';
        div.innerHTML = message.from.username+': '+message.body;
        document.getElementById('messages').appendChild(div);
      });
    });

    url = 'ws://localhost:4000/websocket?token='+token;
    ws = new WebSocket(url);
    ws.onopen = function (event) {
      console.log("Connected!"); 
    };
    ws.onmessage = function (event) {
      var message = JSON.parse(event.data);
      var div = document.createElement('div');
      div.id = 'block';
      div.className = 'block';
      div.innerHTML = message.from+': '+message.body;
      document.getElementById('messages').appendChild(div);
    };

    document.getElementById('message-button').onclick = function(){
      var message = {
        'body': document.getElementById('inputMessage').value,
      };

      var div = document.createElement('div');
      div.id = 'block';
      div.className = 'block';
      div.innerHTML = input.username+': '+message.body;
      document.getElementById('messages').appendChild(div);

      fetch('http://symfony.dev:8007/conversations/1/messages', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
          'Authorization': 'Bearer '+token
        },
        body: JSON.stringify(message)
      });
    }
  });
};
