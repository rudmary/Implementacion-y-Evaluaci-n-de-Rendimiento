var example1 = new Vue({
  el: "#app",
  data: {
    eventos: []
  },
  mounted() {
    axios.get("http://localhost:3001/api/eventos").then(response => {
      console.log("asas");
      this.eventos = response.data;
    });
  }
});
