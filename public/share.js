function drawMainInput() {
    var div = document.createElement('div'),
        input = document.createElement('input');

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

	var keyImage = String.fromCharCode(0x2325)  //

    div.id = 'maincenter';
    div.class = 'simpleForm';

    input.id = 'main';
    input.type = 'password';

	input.placeholder = keyImage;	//

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

function makeItem(archive) {
    var item = document.createElement('li'),
        deleteButton =  document.createElement('input');

    deleteButton.type = 'button';
    deleteButton.value = 'DELETE';

    item.innerHTML = '<span>' + archive.Name + '</archive>';

    deleteButton.addEventListener('click', function (event) {
        // Remove succeded on server
        var li = event.target.parentNode;
        http.delete('/archives/' + archive.Key, function () {
            li.parentNode.removeChild(li);
        });
    });

    item.appendChild(deleteButton); //
    return item;
}

function drawAdmin() {
    var div = document.createElement('div');
    var list = document.createElement('ul');
    var fileInput;

    div.id = 'maincenteradmin';
    div.class = 'admin';

    div.innerHTML = '<h2>File share</h2><h3>Admin menu</h3>';
    div.innerHTML += '<div><h4>Upload archive:</h4><input type="file"></div>';

    fileInput = div.getElementsByTagName('input')[0];

    fileInput.addEventListener('change', function (e) {
        var files = fileInput.files;

        for (var i= 0; i < files.length; i++) {
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

