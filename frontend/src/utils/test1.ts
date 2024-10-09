class WebSocketService {
    private static instance: WebSocketService;
    private socket: WebSocket | null = null;

    private constructor() {}
  
    static getInstance(): WebSocketService {
      if (!WebSocketService.instance) {
        WebSocketService.instance = new WebSocketService();
      }
      return WebSocketService.instance;
    }
  
    connect(url: string) {
      if (!this.socket || this.socket.readyState === WebSocket.CLOSED) {
        this.socket = new WebSocket(url);
  
        this.socket.onopen = () => {
          console.log('WebSocket connection established');
        };
  
        this.socket.onmessage = (event) => {
          console.log('Received message:', event.data);
        };
  
        this.socket.onerror = (error) => {
          console.error('WebSocket error:', error);
        };
  
        this.socket.onclose = () => {
          console.log('WebSocket connection closed');
        };
      }
    }
  
    sendMessage(message: string) {
      if (this.socket && this.socket.readyState === WebSocket.OPEN) {
        this.socket.send(message);
      }
    }
  
    onMessage(callback: (message: string) => void) {
      if (this.socket) {
        this.socket.onmessage = (event) => {
          callback(event.data);
        };
      }
    }
  }
  
//  export default WebSocketService;
  