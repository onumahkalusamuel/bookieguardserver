function ajaxPost(elementId = 'form') {

    var form = document.getElementById(elementId);
    
    var formData = new URLSearchParams(new FormData(form)).toString();

    // create overlay
    var overlay = document.createElement('div');
    overlay.innerHTML = '<div style="display:table;width:100%;height:100vh;position:fixed;top:0;left:0;text-align:center;background-color:#0005;z-index:1000"><div style="display:table-cell;vertical-align:middle;padding-bottom:100px"><span style="color:white;background-color:#118;padding:15px;">please wait...</span></div></div>';

    document.body.appendChild(overlay);

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (xhr.readyState !== 4) return;
        try {
            var response = JSON.parse(xhr.response);
            alert(response.message);
            if (Boolean(response.success) == true && !!response.redirect) {
                window.location.assign(response.redirect);
            }
        } catch (e) {
            alert('An error occured. Please try again later.')
        }

        document.body.removeChild(overlay);
    }

    xhr.open('POST', form.getAttribute('action'), true);
    xhr.setRequestHeader('content-type', 'application/x-www-form-urlencoded');
    xhr.send(formData);
    return false;
}