<template>
<div class="headerContainer">
  <img
      :src="require('@/assets/logo_kz.svg')"
      class="logoKZ"
      alt="Логотип"
  />
  <div class="currTime">
    {{ dateStr }} {{timeStr}}
  </div>
</div>
</template>

<script>
export default {
  name: "HeaderComponent",
  data() {
    return {
      dateStr: "",
      timeStr: "",
    }
  },
  created() {
    setInterval(this.getTime, 1000);
  },
  methods: {
    getTime: function() {
      const today = new Date();
      this.dateStr = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();

      this.timeStr = today.getHours() + ':';
      today.getMinutes().toString().length === 1 ?
          this.timeStr += '0' + today.getMinutes() :
          this.timeStr += today.getMinutes();
      this.timeStr += ':';
      today.getSeconds().toString().length === 1 ?
          this.timeStr += '0' + today.getSeconds() :
          this.timeStr += today.getSeconds() ;
    }
  }
}
</script>

<style scoped>
.headerContainer {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  height: 55px;
  border-bottom: #ce2a2e 2px solid;
  padding: 5px;
  box-shadow: 0 2px 15px 0 #ce2a2e;

}
.logoKZ {
  height:45px;
}
.currTime {
  font-size: 14px;
  color: #ce2a2e;
  margin-right: 10px;
}
</style>