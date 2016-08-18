function isFunction(functionToCheck) {
    var getType = {};
    return functionToCheck && getType.toString.call(functionToCheck) === '[object Function]';
}

var http = {
    jwt: null,
    request: function (method, url, data, success, failure) {
        var self = this;
        var jsonStr = '';
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open(method, url, true)

        xhr.onreadystatechange = function () {
            if (isFunction(success) && xhr.readyState > 3 && xhr.status == 200) {
                success(xhr.responseText);
            } else if (isFunction(failure) && xhr.status >= 400) {
                failure(xhr.status, xhr.responseText)
            }
        };


        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')
        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        if ((method == 'PUT' || method == 'POST') && data) {
            jsonStr = JSON.stringify(data)
            xhr.setRequestHeader('Content-Type', 'application/json')
            xhr.send(jsonStr);
        } else {
            xhr.send();
        }

        return xhr;
    },
    get: function (url, success, failure) {
        var self = this;
        return self.request('GET', url, null, success, failure)
    },
    post: function (url, data, success, failure) {
        var self = this;
        return self.request('POST', url, data, success, failure)
    },
    put: function (url, data, success, failure) {
        var self = this;
        return self.request('POST', url, data, success, failure)
    },
    'delete': function (url, success, failure) {
        var self = this;
        return self.request('DELETE', url, null, success, failure)
    },
    uploadFile: function (url, file, success) {
        var self = this;
        var formData = new FormData();
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open('POST', url, true);
        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')

        formData.append('upload', file, file.name);

        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        xhr.send(formData);
    }
};


function onLogin(res) {
    // Set the JWT
    auth = JSON.parse(res);
    if (auth.token) {
        http.jwt = auth.token;
        drawAdmin();
    }
}

function drawMainInput() {
    var div = document.createElement('div'),
        input = document.createElement('input');

    div.id = 'maincenter';
    div.class = 'simpleForm';

    input.id = 'main';
    input.type = 'password';

    input.addEventListener('keypress', function (e) {
        var key = window.event ? e.keyCode : e.which;
        var colon = input.value.indexOf(':')
        var creds = [];

        if (key == 13) {
            if (colon < 0) {
                http.get('/archive/' + input.value);
            } else {
                creds = input.value.split(':', 2)
                http.post('/login', {'name': creds[0], 'password': creds[1]}, onLogin);
            }
        }
    });

    div.appendChild(input);

    document.body.innerHTML = '';
    document.body.appendChild(div);
}

function drawAdmin() {
    var div = document.createElement('div');
    var list = document.createElement('ul');
    var archives = [];

    div.id = 'maincenter';
    div.class = 'admin';

    div.innerHTML = '<h2>File share</h2><h3>Admin menu</h3>';
    div.innerHTML += '<div><h4>Upload archive:</h4><input type="file"></div>';

    http.get('/archives', function (r) {
        var li;
        resp = JSON.parse(r);

        if (!resp.archives || resp.archives.length == 0) {
            return;
        }

        for (archive in archives) {
            li = document.createElement('li');
            li.innerHTML = '<span>' + archive.Name + '</span>';
            list.appendChild(li);
        }
        div.appendChild(list);
    });

    document.body.innerHTML = '';
    document.body.appendChild(div);
}

document.addEventListener('DOMContentLoaded', drawMainInput);
