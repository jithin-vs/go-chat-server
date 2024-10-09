class WebSocketService {
  private socket: WebSocket | null = null;
  private messageCallbacks: Array<(message: string) => void> = []; 

  constructor(url: string) {
    this.connect(url);
  }

  connect(url: string) {
    if (!this.socket || this.socket.readyState === WebSocket.CLOSED) {
      this.socket = new WebSocket(url);
      this.socket.onopen = () => {
        console.log("WebSocket connection established on ", url);
      };
      this.socket.onmessage = (e) => {
        const data = JSON.parse(e.data);
        console.log("Received Message from", url, "message:", data.body);
        this.messageCallbacks.forEach(callback => callback(data.body)); 
      };
      this.socket.onerror = () => {
        console.log("Error connecting to WebSocket ");
      };
      this.socket.onclose = () => {
        console.log("WebSocket connection closed on ", url);
      };
    }
  }
    
    sendMessage(message: string) {
      if (this.socket && this.socket.readyState == WebSocket.OPEN) {
            console.log(message);
            this.socket.send(message);
        }
    }

    onMessage(callback: (message: string) => void) {
      this.messageCallbacks.push(callback); // Add the callback to the array
    }

    close() {
        if (this.socket) {
            this.socket.close();
        }
    }
}

export default WebSocketService;
