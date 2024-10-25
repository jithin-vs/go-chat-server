class WebSocketService {

  private socket: WebSocket | null = null;
  private url : string ;
  private messageCallbacks: Array<(message: string) => void> = [];
  private messageQueue: string[] =  [];
  private reconectInterval: number = 5000;
  private manualClose : boolean = false;

  constructor(url: string) {
    this.url = url;
    this.connect(url);
  }

  connect(url: string) {
    if (!this.socket || this.socket.readyState === WebSocket.CLOSED) {
      this.socket = new WebSocket(url);
      this.manualClose = false;
      this.socket.onopen = () => {
        console.log("WebSocket connection established on ", url);

        while (this.messageQueue.length > 0) {
          const message = this.messageQueue.shift();
          if (message) this.sendMessage(message);
        }
      };
      this.socket.onmessage = (e) => {
        const data = JSON.parse(e.data);
        console.log("Received Message from", url, "message:", e.data);
        this.messageCallbacks.forEach(callback => callback(e.data)); 
      };
      this.socket.onerror = () => {
        console.log("Error connecting to WebSocket ");
      };
      this.socket.onclose = () => {
        console.log("WebSocket connection closed on ", url);
        if (!this.manualClose) {
          this.reconnect
        }
      };
    }
  }
    
    sendMessage(message: string) {
      if (this.socket && this.socket.readyState == WebSocket.OPEN) {
            console.log(message);
            this.socket.send(message);
      } else {
            this.messageQueue.push(message);
        }
    }

    onMessage(callback: (message: string) => void) {
      this.messageCallbacks.push(callback); 
    }
  reconnect() { 
    setTimeout(() => {
      this.connect(this.url);
    }, this.reconectInterval); 
  }
   
  close() {
       this.manualClose = true;
        if (this.socket) {
            this.socket.close();
        }
    }
}

export default WebSocketService;
