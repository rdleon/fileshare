function drawMainInput() {
    var div = document.createElement('div'),
        input = document.createElement('input'),
        keyImage = String.fromCharCode(0x2325);

    function onLogin(res) {
        // Set the JWT
        if (res.token) {
            http.setJWT(res.token);
            drawAdmin();
        }
    }

    if (http.getJWT()) {
        drawAdmin();
        return;
    }

    div.id = 'maincenter';
    div.class = 'simpleForm';

    input.id = 'main';
    input.type = 'password';

	input.placeholder = keyImage;

    input.addEventListener('keypress', function (e) {
        var key = window.event ? e.keyCode : e.which,
            colon = input.value.indexOf(':'),
            creds = [];

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

function makeItem(archive) {
    var item = document.createElement('li'),
        deleteButton = document.createElement('input'),
        expiration = new Date(archive.Expire * 1000);

    item.innerHTML = '<span>';
    item.innerHTML += '<span>';
    item.innerHTML += '<a href="/archives/'+ archive.Key +'">' + archive.Name + '</a>'
    item.innerHTML += '<span> Expires: '+ expiration.toISOString() +'</span>';
    item.innerHTML += '</span>';

    deleteButton.type = 'button';
    deleteButton.value = 'DELETE';

    deleteButton.addEventListener('click', function (event) {
        // Remove if deleted succeeded on server
        http.delete('/archives/' + archive.Key, function () {
            item.parentNode.removeChild(item);
        });
    });

    item.appendChild(deleteButton);

    return item;
}

function drawAdmin() {
    var div = document.createElement('div'),
        list = document.createElement('ul'),
        fileInput;

    div.id = 'maincenteradmin';
    div.class = 'admin';

    div.innerHTML = '<h2>File share</h2><h3>Admin menu</h3>';
    div.innerHTML += '<div><h4>Upload archive:</h4><input type="file"></div>';

    fileInput = div.getElementsByTagName('input')[0];

    fileInput.addEventListener('change', function (e) {
        var i, files = fileInput.files;

        for (i = 0; i < files.length; i++) {
            http.uploadFile('/archives', files[i], function (newArchive) {
                li = makeItem(newArchive)
                list.appendChild(li);
            });
        }
    });

    http.get('/archives', function (resp) {
        var li, i, deleteButton;

        if (!resp.archives || resp.archives.length == 0) {
            // No archives to show
            return;
        }

        for (i in resp.archives) {
            li = makeItem(resp.archives[i])
            list.appendChild(li);
        }
    });

    div.appendChild(list);

    document.body.innerHTML = '';
    document.body.appendChild(div);
}

var http = new HTTP();
document.addEventListener('DOMContentLoaded', drawMainInput);
