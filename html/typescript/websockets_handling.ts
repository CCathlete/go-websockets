declare const notie: any;

let socket: WebSocket | null = null;
const output: HTMLElement | null = document.getElementById("output");
const userField: HTMLInputElement | null = document.getElementById("username") as HTMLInputElement;
const messageField: HTMLInputElement | null = document.getElementById("message") as HTMLInputElement;

window.onbeforeunload = () => {
    console.log("Leaving.");
    if (socket) {
        const jsonData: { action: string } = { action: "left" };
        socket.send(JSON.stringify(jsonData));
    }
};

document.addEventListener("DOMContentLoaded", () => {
    socket = new WebSocket("ws://127.0.0.1:8080/ws");

    socket.onopen = () => {
        console.log("Successfully connected via websocket.");
    };

    socket.onclose = () => {
        console.log("Connection closed.");
    };

    socket.onerror = (error: Event) => {
        console.log("There was an error: ", error);
    };

    socket.onmessage = (msg: MessageEvent) => {
        const j: any = JSON.parse(msg.data);
        console.log("Action is: ", j.action);

        switch (j.action) {
            case "list_users":
                const ul: HTMLElement | null = document.getElementById("online_users");
                if (ul) {
                    while (ul.firstChild) {
                        ul.removeChild(ul.firstChild);
                    }

                    if (j.connected_users.length > 0) {
                        j.connected_users.forEach((element: string) => {
                            const li: HTMLLIElement = document.createElement("li");
                            li.appendChild(document.createTextNode(element));
                            ul.appendChild(li);
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

    userField?.addEventListener("change", function () {
        if (socket) {
            const jsonData: { action: string; username: string } = {
                action: "username",
                username: this.value,
            };
            socket.send(JSON.stringify(jsonData));
        }
    });

    document.getElementById("message")?.addEventListener("keydown", function (event: KeyboardEvent) {
        if (event.code === "Enter") {
            if (!socket) {
                console.log("No connection.");
                return false;
            }
            if (!userField?.value || !messageField?.value) {
                errorMessage("Please enter a username and message.");
                return false;
            }
            event.preventDefault();
            event.stopPropagation();
            sendMessage();
        }
    });

    document.getElementById("sendBtn")?.addEventListener("click", function () {
        if (!userField?.value || !messageField?.value) {
            errorMessage("Please enter a username and message.");
            return false;
        }

        sendMessage();
    });
});

function sendMessage() {
    if (!userField || !messageField || !socket) return;

    const jsonData: { action: string; username: string; message: string } = {
        action: "broadcast",
        username: userField.value,
        message: messageField.value,
    };
    socket.send(JSON.stringify(jsonData));
    messageField.value = "";
}

function errorMessage(msg: string) {
    notie.alert({
        type: 'error',
        text: msg,
    });
}
