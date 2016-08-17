var http = {
    jwt: null,
    request: function (method, url, data, success, failure) {
        var self = this;
        var jsonStr = '';
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open(method, url, true)

        xhr.onreadystatechange = function () {
            if (xhr.readyState > 3 && xhr.status = 200) {
                success(xhr.responseText);
            } else if (xhr.status >= 400) {
                failure(xhr.status)
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
        return request('GET', url, null, success, failure)
    },
    post: function (url, data, success, failure) {
        return request('POST', url, data, success, failure)
    },
    put: function (url, data, success, failure) {
        return request('POST', url, data, success, failure)
    },
    'delete': function (url, success, failure) {
        return request('DELETE', url, null, success, failure)
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

window.onload = function () {
};
