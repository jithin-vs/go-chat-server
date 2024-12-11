class WebSocketService {
  private socket: WebSocket | null = null;
  private url: string;
  private messageCallbacks: Array<(message: string) => void> = [];
  private messageQueue: string[] = [];
  private reconectInterval: number = 5000;
  private manualClose: boolean = false;

  constructor(url: string) {
    this.url = url;
    console.log("url",this.url);
    this.connect(url);
  }

  connect(url: string) {
    if (!this.socket || this.socket.readyState === WebSocket.CLOSED) {
      try {
        this.socket = new WebSocket(url);
        this.manualClose = false;

        this.socket.onopen = () => {
          console.log(`[WebSocket] Connection established on: ${url}`);

          // Process queued messages
          while (this.messageQueue.length > 0) {
            const message = this.messageQueue.shift();
            if (message) this.sendMessage(message);
          }
        };

        this.socket.onmessage = (e) => {
          console.log(`[WebSocket] Message received from ${url}:`, e.data);
          try {
            const data = JSON.parse(e.data);
            console.log(`[WebSocket] Parsed message:`, data);
            this.messageCallbacks.forEach((callback) => callback(data));
          } catch (parseError) {
            console.error(
              `[WebSocket] Failed to parse message:`,
              e.data,
              parseError
            );
          }
        };

        this.socket.onerror = (error) => {
          console.error(
            `[WebSocket] Error occurred while connecting to ${url}:`,
            error
          );
        };

        this.socket.onclose = (event) => {
          console.warn(
            `[WebSocket] Connection closed on ${url}:`,
            event.reason || "No reason provided"
          );
          if (!this.manualClose) {
            console.log(`[WebSocket] Attempting to reconnect to ${url}...`);
            this.reconnect();
          }
        };
      } catch (connectionError) {
        console.error(
          `[WebSocket] Failed to establish connection to ${url}:`,
          connectionError
        );
      }
    } else {
      console.log(
        `[WebSocket] Connection attempt skipped as the socket is already open or in progress.`
      );
    }
  }

  isOpen(): boolean {
    return this.socket?.readyState === WebSocket.OPEN;
  }

  sendMessage(message: string) {
    if (this.socket && this.socket.readyState == WebSocket.OPEN) {
      console.log(message);
      this.socket.send(message);
    } else {
      this.messageQueue.push(message);
    }
  }
  emit(event: string, data: any) {
    const payload = JSON.stringify({ event, data });
    this.sendMessage(payload);
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
