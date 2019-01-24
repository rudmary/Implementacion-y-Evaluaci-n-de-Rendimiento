var example1 = new Vue({
  el: "#app",
  data: {
    eventos: [],
    asientos: [],
    evento_id: -1
  },
  mounted() {
    $(".modal").modal();
    axios.get("http://localhost:3000/api/eventos").then(response => {
      this.eventos = response.data;
    });
  },
  methods: {
    async obtenerAsientos(localidadId, evento_id) {
      const { data } = await axios.get(
        "http://localhost:3000/api/asientos/" + localidadId
      );
      this.asientos = data.filter(asiento => {
        if (asiento) {
          return asiento;
        }
      });
      this.evento_id = evento_id;
    },
    async comprarBoleto(asiento_id, usuario_id) {
      console.log({ asiento_id, usuario_id, evento_id: this.evento_id });
    }
  }
});
