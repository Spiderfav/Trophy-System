var loginBox = document.getElementById("login");
var regBox = document.getElementById("register");
var forgetBox = document.getElementById("forgot");

var loginTab = document.getElementById("lt");
var regTab = document.getElementById("rt");

function regTabFun() {
    event.preventDefault();

    regBox.style.visibility = "visible";
    loginBox.style.visibility = "hidden";
    forgetBox.style.visibility = "hidden";

    regTab.style.backgroundColor = "rgb(12, 132, 189)";
    loginTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
}
function loginTabFun() {
    event.preventDefault();

    regBox.style.visibility = "hidden";
    loginBox.style.visibility = "visible";
    forgetBox.style.visibility = "hidden";

    loginTab.style.backgroundColor = "rgb(12, 132, 189)";
    regTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
}
function forTabFun() {
    event.preventDefault();

    regBox.style.visibility = "hidden";
    loginBox.style.visibility = "hidden";
    forgetBox.style.visibility = "visible";

    regTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";
    loginTab.style.backgroundColor = "rgba(11, 177, 224, 0.82)";

}

$(document).ready(function () {
    $("#register_form").submit(function () {

        var username = document.getElementById("us").value;
        var email = document.getElementById("re").value;
        var password = document.getElementById("rp").value;
        var passwordRetype = document.getElementById("rrp").value;

        if (email == "") {
            alert("Email required.");
            return;
        }
        else if (password == "") {
            alert("Password required.");
            return;
        }
        else if (username == "") {
            alert("Username required.");
            return;
        }
        else if (passwordRetype == "") {
            alert("Password required.");
            return;
        }
        else if (password != passwordRetype) {
            alert("Passwords don't match, retype your Password.");
            return;
        }
        else {
            $.ajax({
                url: '/createuser',
                type: 'post',
                data: { username: username, password: password, email: email },
                success: function (response) {
                    var msg = "";
                    if (response == true) {
                        window.location = "/homepage";  // window.location = "/homepage";
                    } else {
                        msg = "Invalid username and password!";
                    }
                    $("#message").html(msg);
                    console.log(response);

                    alert(username + "  Thanks for registration. \nTry to login Now");

                    document.getElementById("re").value = "";
                    document.getElementById("rp").value = "";
                    document.getElementById("rrp").value = "";
                    document.getElementById("us").value = "";


                }
            })
        }
    });
    $("#login_form").submit(function () {
        var email = document.getElementById("se").value;
        var password = document.getElementById("sp").value;

        if (email == "") {
            alert("Email required.");
            return;
        }
        else if (password == "") {
            alert("Password required.");
            return;
        }
        else {
            var success
            $.ajax({
                url: '/userlogin',
                type: 'post',
                data: { email: email, password: password },
                dataType: 'json',
                async: false,
                success: function (response) {
                    success = response;
                }
            });
            var obj = jQuery.parseJSON(success)
            var msg = "";
            alert(obj.status === "true");

            if (obj.status === true) {
                window.location = "/homepage";  // window.location = "/homepage";
            } else {
                msg = "Invalid username and password!";
            }
            $("#message").html(msg);
            console.log(success);

            alert(email + ", Login successful!");

            document.getElementById("se").value = "";
            document.getElementById("sp").value = "";

        }
    });
});



function forgot() {
    event.preventDefault();

    var email = document.getElementById("fe").value;

    if (emailArray.indexOf(email) == -1) {
        if (email == "") {
            alert("Email required.");
            return;
        }
        alert("Email does not exist.");
        return;
    }

    alert("email is send to your email check it in 24hr. \n Thanks");
    document.getElementById("fe").value = "";
}