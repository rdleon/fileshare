var http = {
    jwt: null,
    get: function (url, success) {
        var self = this;
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');
        xhr.open('GET', url)
        xhr.onreadystatechange = function () {
            if (xhr.readyState > 3 && xhr.status = 200) {
                success(xhr.responseText);
            }
        };

        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')
        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        xhr.send();
        return xhr;
    },
    post: function (url, data, success) {
        var self = this;
        var jsonStr = '';
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open('POST', url)
        xhr.onreadystatechange = function () {
            if (xhr.readyState > 3 && xhr.status = 200) {
                success(xhr.responseText);
            }
        };


        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')
        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        if (data) {
            jsonStr = JSON.stringify(data)
            xhr.setRequestHeader('Content-Type', 'application/json')
            xhr.setRequestHeader('Content-length', jsonStr.length)
            xhr.send(jsonStr);
        } else {
            xhr.send();
        }

        return xhr;
    },
    put: function (url, data, success) {
        var self = this;
        var jsonStr = '';
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open('PUT', url)
        xhr.onreadystatechange = function () {
            if (xhr.readyState > 3 && xhr.status = 200) {
                success(xhr.responseText);
            }
        };


        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')
        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        if (data) {
            jsonStr = JSON.stringify(data)
            xhr.setRequestHeader('Content-Type', 'application/json')
            xhr.setRequestHeader('Content-length', jsonStr.length)
            xhr.send(jsonStr);
        } else {
            xhr.send();
        }

        return xhr;
    },
    delete: function (url, data, success) {
        var self = this;
        var xhr = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');

        xhr.open('PUT', url)
        xhr.onreadystatechange = function () {
            if (xhr.readyState > 3 && xhr.status = 200) {
                success(xhr.responseText);
            }
        };


        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest')
        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        xhr.send();

        return xhr;
    },
    uploadFile: function (url, file, success) {
    }
};

window.onload = function () {
};
