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
            xhr.setRequestHeader('Content-length', jsonStr.length)
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

function drawMainInput() {
    var div = document.createElement('div'),
        input = document.createElement('input');

    input.id = 'main';
    input.type = 'password';

    input.addEventListener('keypress', function (e) {
        var key = window.event ? e.keyCode : e.which;
        if (key == 13) {
            http.get('/archive/' + input.value);
        }
    });

    div.appendChild(input);

    document.body.innerHTML = '';
    document.body.appendChild(div);
}

function drawAdmin() {
}

document.addEventListener('DOMContentLoaded', drawMainInput);
