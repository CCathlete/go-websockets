var socket = null;
var output = document.getElementById("output");
var userField = document.getElementById("username");
var messageField = document.getElementById("message");
window.onbeforeunload = function () {
    console.log("Leaving.");
    if (socket) {
        var jsonData = { action: "left" };
        socket.send(JSON.stringify(jsonData));
    }
};
document.addEventListener("DOMContentLoaded", function () {
    var _a, _b;
    socket = new WebSocket("ws://127.0.0.1:8080/ws");
    socket.onopen = function () {
        console.log("Successfully connected via websocket.");
    };
    socket.onclose = function () {
        console.log("Connection closed.");
    };
    socket.onerror = function (error) {
        console.log("There was an error: ", error);
    };
    socket.onmessage = function (msg) {
        var j = JSON.parse(msg.data);
        console.log("Action is: ", j.action);
        switch (j.action) {
            case "list_users":
                var ul_1 = document.getElementById("online_users");
                if (ul_1) {
                    while (ul_1.firstChild) {
                        ul_1.removeChild(ul_1.firstChild);
                    }
                    if (j.connected_users.length > 0) {
                        j.connected_users.forEach(function (element) {
                            var li = document.createElement("li");
                            li.appendChild(document.createTextNode(element));
                            ul_1.appendChild(li);
                        });
                    }
                }
                break;
            case "broadcast":
                if (output) {
                    output.innerHTML += j.message + "<br>";
                }
                break;
        }
    };
    userField === null || userField === void 0 ? void 0 : userField.addEventListener("change", function () {
        if (socket) {
            var jsonData = {
                action: "username",
                username: this.value,
            };
            socket.send(JSON.stringify(jsonData));
        }
    });
    (_a = document.getElementById("message")) === null || _a === void 0 ? void 0 : _a.addEventListener("keydown", function (event) {
        if (event.code === "Enter") {
            if (!socket) {
                console.log("No connection.");
                return false;
            }
            if (!(userField === null || userField === void 0 ? void 0 : userField.value) || !(messageField === null || messageField === void 0 ? void 0 : messageField.value)) {
                errorMessage("Please enter a username and message.");
                return false;
            }
            event.preventDefault();
            event.stopPropagation();
            sendMessage();
        }
    });
    (_b = document.getElementById("sendBtn")) === null || _b === void 0 ? void 0 : _b.addEventListener("click", function () {
        if (!(userField === null || userField === void 0 ? void 0 : userField.value) || !(messageField === null || messageField === void 0 ? void 0 : messageField.value)) {
            errorMessage("Please enter a username and message.");
            return false;
        }
        sendMessage();
    });
});
function sendMessage() {
    if (!userField || !messageField || !socket)
        return;
    var jsonData = {
        action: "broadcast",
        username: userField.value,
        message: messageField.value,
    };
    socket.send(JSON.stringify(jsonData));
    messageField.value = "";
}
function errorMessage(msg) {
    notie.alert({
        type: 'error',
        text: msg,
    });
}
