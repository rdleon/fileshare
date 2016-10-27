"use strict";

function isFunction(func) {
    var type = {};
    return func && type.toString.call(func) === '[object Function]';
}

function HTTP() {
    var self = this;

    self.jwt = localStorage.getItem('webtoken');

    self.checkToken();
};

HTTP.prototype.request = function (method, url, body, success, failure) {
        var self = this,
            jsonStr,
            xhr;

        if (window.XMLHttpRequest) {
            xhr = new XMLHttpRequest();
        } else {
            xhr = new ActiveXObject('Microsoft.XMLHTTP');
        }

        xhr.open(method, url, true);

        xhr.onreadystatechange = function () {
            var resp = null;

            if (isFunction(success) && xhr.readyState > 3 && xhr.status == 200) {
                if (xhr.responseText) {
                    resp = JSON.parse(xhr.responseText);
                    success(resp);
                } else {
                    success({});
                }

            } else if (isFunction(failure) && xhr.status >= 400) {
                if (xhr.responseText) {
                    resp = JSON.parse(xhr.responseText);
                    failure(xhr.status, resp);
                } else {
                    failure(xhr.status, {});
                }
            }
        };

        xhr.setRequestHeader('Accept', 'application/json');
        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest');

        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        if ((method == 'PUT' || method == 'POST') && body) {
            jsonStr = JSON.stringify(body);
            xhr.setRequestHeader('Content-Type', 'application/json');
            xhr.send(jsonStr);
        } else {
            xhr.send();
        }

        return xhr;
};

HTTP.prototype.setJWT = function (jwt) {
    this.jwt = jwt;
    localStorage.setItem('webtoken', jwt);
};

HTTP.prototype.getJWT = function (jwt) {
    if (!this.jwt) {
        this.jwt = localStorage.getItem('webtoken');
        this.checkToken();
    }

    return this.jwt;
};

HTTP.prototype.clearJWT = function () {
    this.jwt = null;
    localStorage.removeItem('webtoken');
};

HTTP.prototype.get = function(url, success, failure) {
    return this.request('GET', url, null, success, failure);
};

HTTP.prototype.post = function(url, body, success, failure) {
    return this.request('POST', url, body, success, failure);
};

HTTP.prototype.put = function(url, body, success, failure) {
    return this.request('PUT', url, body, success, failure);
};

HTTP.prototype.delete = function(url) {
    return this.request('DELETE', url, null, success, failure);
};

HTTP.prototype.uploadFile = function(url, file, success, failure) {
        var self = this,
            xhr,
            formData;

        if (window.XMLHttpRequest) {
            xhr = new XMLHttpRequest();
        } else {
            xhr = new ActiveXObject('Microsoft.XMLHTTP');
        }

        xhr.open('POST', url, true);
        xhr.setRequestHeader('Accept', 'application/json');
        xhr.setRequestHeader('X-Request-With', 'XMLHttpRequest');

        xhr.onreadystatechange = function () {
            var resp = null;

            if (isFunction(success) && xhr.readyState > 3 && xhr.status == 200) {
                if (xhr.responseText) {
                    resp = JSON.parse(xhr.responseText);
                }
                success(resp);
            } else if (isFunction(failure) && xhr.status >= 400) {
                if (xhr.responseText) {
                    resp = JSON.parse(xhr.responseText);
                }
                failure(xhr.status, resp);
            }
        };

        formData = new FormData();
        formData.append('upload', file, file.name);

        if (self.jwt) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + self.jwt);
        }

        xhr.send(formData);
};

HTTP.prototype.checkToken = function() {
    var valid = true,
        token = [],
        user = {},
        now = new Date();

    if (this.jwt) {
        token = this.jwt.split('.');
        if (token.length > 1) {
            user = JSON.parse(atob(token[1]));
            if (!user.exp || user.exp > now.getTime()) {
                valid = false;
                this.jwt = null;
            }
        } else {
            valid = false;
        }
    } else {
        valid = false;
    }

    return valid;
};
