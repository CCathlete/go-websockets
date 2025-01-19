// Stating that  notie is being loaded and defined somewhere else (before calling this script).
declare const notie: any;
declare class ReconnectingWebSocket {
    constructor(
        url: string,
        protocols?: null | string | string[],
        options?: any
    );
    send(data: string): void;
    close(): void;
    onopen?: () => void;
    onclose?: () => void;
    onmessage?: (event: MessageEvent) => void;
    onerror?: (error: Event) => void;
}


let socket: ReconnectingWebSocket | null = null;
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
    socket = new ReconnectingWebSocket(
        "ws://127.0.0.1:8080/ws",
        null,
        { debug: true, reconnectInterval: 3000 }
    );

    const offline: string =
        `<span class="badge bg-danger">Offline</span>`;
    const online: string =
        `<span class="badge bg-success">Connected</span>`;
    let statusDiv: HTMLElement | null =
        document.getElementById("status");

    socket.onopen = () => {
        console.log("Successfully connected via websocket.");
        if (statusDiv) {
            statusDiv.innerHTML = online;
        };
    };

    socket.onclose = () => {
        console.log("Connection closed.");
        if (statusDiv) {
            statusDiv.innerHTML = offline;
        };
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
