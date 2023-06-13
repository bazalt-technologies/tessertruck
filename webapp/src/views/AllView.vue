<template>
<div class="allContainer">
  <!-- Левая панель с тракторами -->
  <div class="tracsList">
    <div class="upper-r-table">
      <div class="title">
        Тракторы:
      </div>
      <div>
        <b-button variant="danger" class="upper-r-table-button" @click="showAddTracModal = true"><b-icon-plus-circle/></b-button>
      </div>
    </div>
    <div>
      <ag-grid-vue
          class="ag-theme-balham"
          :columnDefs="columns"
          :rowData="tractors"
          :rowSelection="rowSelection"
          style="width: 682px"
          @selection-changed="onSelectionChanged()"
          @grid-ready="onGridReady"
      />
    </div>
  </div>

  <!-- Правая панель с информацией по трактору -->
  <div class="tracInfo" v-if="selectedTracID">
    <div class="upper-r-table">

      <div class="title" v-if="browserWidth > 1380">
        Трактор №{{selectedTracID}} - {{selectedTracName}}:
      </div>
      <div class="upper-r-table-btns">
        <b-button
            variant="danger" class="upper-r-table-button"
            @click="showAddRuleModal = true"
            style="margin-left: 0"
        >
          <b-icon-plus-circle/> Добавить правило
        </b-button>
        <b-button variant="danger" class="upper-r-table-button" @click="downloadJournal">
          <b-icon-journal-text/> Журнал нарушений
        </b-button>

      </div>

    </div>
    <div>
      <div class="infoField" v-for="item in columnsInfo" :key="item.field">
        {{item.label}}: {{ rowsInfo[item.field] }} {{item.measure}}
      </div>
    </div>
  </div>




  <!-- Модалка для добавления правила -->
  <b-modal
    id="add-rule-modal"
    v-model="showAddRuleModal"
    :title="'Добавить правило'"
    header-bg-variant="danger"
    header-text-variant="light"
    header-border-variant="danger"
    footer-border-variant="danger"
    ok-variant="success"
    ok-title="Добавить"
    cancel-title="Отменить"
    hide-header-close
    @ok="addRule"
  >
    <b-alert variant="danger" :show="badInput">Все поля должны быть заполнены</b-alert>
    <b-form-input
      placeholder="Название"
      v-model="newRule.Name"
      class="inputField"
    />
    <b-form-select
      v-model="newRule.FieldName"
      :options="selectField"
      class="inputField"
    />
    <b-form-input
      placeholder="Значение"
      v-model="newRule.Val"
      class="inputField"
    />
  </b-modal>

<!--  Модалка для добавления трактора-->
  <b-modal
      id="add-trac-modal"
      v-model="showAddTracModal"
      :title="'Добавить трактор'"
      header-bg-variant="danger"
      header-text-variant="light"
      header-border-variant="danger"
      footer-border-variant="danger"
      ok-variant="success"
      ok-title="Добавить"
      cancel-title="Отменить"
      hide-header-close
      @ok="addTrac"
  >
    <b-alert variant="danger" :show="badInput">Все поля должны быть заполнены</b-alert>
    <b-form-input
        placeholder="Название"
        v-model="tracName"
        class="inputField"
    />
    <b-form-datepicker
        placeholder="Дата создания"
        v-model="createDate"
        class="inputField"
        locale="ru"
    />
    <b-form-datepicker
      placeholder="В эксплуатации с"
      v-model="useDate"
      class="inputField"
      locale="ru"
    />
    <b-form-input
        placeholder="Место эксплуатации"
        v-model="usePlace"
        class="inputField"
    />
  </b-modal>
</div>
</template>

<script>
import "ag-grid-community/styles/ag-grid.css";
import "ag-grid-community/styles/ag-theme-balham.css";
import {AgGridVue} from "ag-grid-vue";
//import {url} from "@/main.js"
export default {
  name: "AllView",
  components: {
    AgGridVue
  },
  data() {return {
    showAddRuleModal: false,
    showAddTracModal: false,
    connection: null,
    badInput: false,
    newRule: {
      id: 0,
      Name: "",
      TractorId: this.selectedTracID,
      FieldName: null,
      ValInt: "",
      ValFloat: "",
      Val: null,
    },
    newTrac: {},
    rowSelection: 'single',
    selectedTracID: 0,
    selectedTracName: '',
    columns: [
      {field: "ID", headerName: 'Номер', sortable: true, width: 80},
      {field: "Name", headerName: 'Название', sortable: true, width: 150},
      {field: "CreateDate", headerName: 'Дата создания', sortable: true, width: 150},
      {field: "UseDate", headerName: 'В эксплуатации с', sortable: true, width: 150},
      {field: "UsePlace", headerName: 'Место эксплуатации', sortable: true, width: 150},

    ],
    tractors: [],
    selectField: [
      {value: null, text: "Выберите поле"},
      {value: "SpeedRT", text: "Скорость"},
      {value: "EngineRPS", text: "Оборотов в секунду"},
      {value: "FuelLevel", text: "Уровень топлива"},
      {value: "FuelSpent", text: "Потрачено топлива"},
      {value: "FuelConsumption", text: "Расход топлива"},
      {value: "EngineRegime", text: "Режим двигателя"},
      {value: "OilTemperature", text: "Температура масла"},
      {value: "EnvTemperature", text: "Внешняя температура"},
      {value: "RedLamp", text: "Красный индикатор"},
      {value: "WarnLamp", text: "Жёлтый индикатор"},
    ],
    columnsInfo: [
      {field: "SpeedRT", label: "Скорость", measure: "км/ч" },
      {field: "EngineRPS", label: "Оборотов в секунду", measure: "" },
      {field: "FuelLevel", label: "Уровень топлива", measure: "л" },
      {field: "FuelSpent", label: "Потрачено топлива", measure: "л" },
      {field: "FuelConsumption", label: "Расход топлива", measure: "л" },
      {field: "EngineRegime", label: "Режим двигателя", measure: "" },
      {field: "OilTemperature", label: "Температура масла", measure: "°С" },
      {field: "EnvTemperature", label: "Внешняя температура", measure: "°С" },
      {field: "RedLamp", label: "Красный индикатор", measure: "" },
      {field: "WarnLamp", label: "Жёлтый индикатор", measure: "" },
    ],
    wsData: {
      ID: null,
      Name: "",
      Teledata: {
        SpeedRT: [],
        EngineRPS: [],
        FuelLevel: [],
        FuelSpent: [],
        FuelConsumption: [],
        EngineRegime: [],
        OilTemperature: [],
        EnvTemperature: [],
        RedLamp: [],
        WarnLamp: [],
        RulesWarnings: [],
      }
    },
    rowsInfo: {
      SpeedRT: 0,
      EngineRPS: 0,
      FuelLevel: 0,
      FuelSpent: 0,
      FuelConsumption: 0,
      EngineRegime: 0,
      OilTemperature: 0,
      EnvTemperature: 0,
      RedLamp: 0,
      WarnLamp: 0,
    },
    browserWidth: 0,
    browserHeight: 0,
  }},
  computed: {
    createDate: {
      get() {
        if (!this.newTrac) {
          return ""
        }
        return this.newTrac.CreateDate
      },
      set(val) {
        this.newTrac.CreateDate = val
      }
    },
    useDate: {
      get() {
        if (!this.newTrac) {
          return ""
        }
        return this.newTrac.UseDate
      },
      set(val) {
        this.newTrac.UseDate = val
      }
    },
    tracName: {
      get() {
        if (!this.newTrac) {
          return ""
        }
        return this.newTrac.Name
      },
      set(val) {
        this.newTrac.Name = val
      }
    },
    usePlace: {
      get() {
        if (!this.newTrac) {
          return ""
        }
        return  this.newTrac.UsePlace
      },
      set(val) {
        this.newTrac.UsePlace = val
      }
    }
  },
  created() {
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
    this.getData()

  },
  onUnmounted() {
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    getData() {
      this.$http.get("http://localhost:8085/api/v1/tractors").then(
          // eslint-disable-next-line no-unused-vars
          response=> {
            this.tractors = response && response.data ? response.data : []
            // eslint-disable-next-line no-unused-vars
          })
    },
    isIn(arr, str) {
      for (let i = 0; i < arr.length; i++) {
        if (str === arr[i]) {
          return true
        }
      }
      return false
    },
    setInfo(newData) {
      let someData = JSON.parse(newData)
      this.rowsInfo.SpeedRT = someData.Teledata.SpeedRT
      this.rowsInfo.EngineRPS = someData.Teledata.EngineRPS
      this.rowsInfo.FuelLevel = someData.Teledata.FuelLevel
      this.rowsInfo.FuelSpent = someData.Teledata.FuelSpent
      this.rowsInfo.FuelConsumption = someData.Teledata.FuelConsumption
      this.rowsInfo.EngineRegime = someData.Teledata.EngineRegime
      this.rowsInfo.OilTemperature = someData.Teledata.OilTemperature
      this.rowsInfo.EnvTemperature = someData.Teledata.EnvTemperature
      this.rowsInfo.RedLamp = someData.Teledata.RedLamp
      this.rowsInfo.WarnLamp = someData.Teledata.WarnLamp
    },
    isNumber(str) {
      if (str === null || str.length === 0) {
        return false
      }
      for (let i = 0; i < str.length; i++) {
        if(str[i] > '9' || str[i] < '0') {
          if (str[i] !== '.') {
            return false;
          }
        }
      }
      return true;
    },
    checkNewRuleInput() {
      return this.newRule.Name !== "" &&
          this.newRule.FieldName !== null &&
          this.newRule.ValInt !== null &&
          this.isNumber(this.newRule.Val)
    },
    addTrac(bvModalEvent) {
      bvModalEvent.preventDefault()
      this.showAddTracModal = false
      this.$http.post("http://localhost:8085/api/v1/tractors", this.newTrac).then(
          // eslint-disable-next-line no-unused-vars
          response=> {
            this.newTrac = {}
            this.getData()
            // eslint-disable-next-line no-unused-vars
          }
    )

    },
    addRule(bvModalEvent) {
      bvModalEvent.preventDefault()
      if (!this.checkNewRuleInput()) {
        this.badInput = true
        return
      }
      this.badInput = false
      switch (this.newRule.FieldName) {
        case 'SpeedRT':case'FuelLevel':case'FuelSpent':
        case 'FuelConsumption':case 'OilTemperature':
        case 'EnvTemperature':
          this.newRule.ValFloat = Number(this.newRule.Val);
          break;
        default:
          this.newRule.ValInt = Math.floor(Number(this.newRule.Val));
          break;
      }
      let sending = {
        ID: 0,
        Name: this.newRule.Name,
        TractorID: this.selectedTracID,
        FieldName: this.newRule.FieldName,
        ValInt: this.newRule.ValInt || 0,
        ValFloat: this.newRule.ValFloat || 0,
      }
      this.showAddRuleModal = false
      this.newRule.Name = "";
      this.newRule.ValInt = 0;
      this.newRule.ValFloat = 0;
      this.newRule.FieldName = null;
      this.newRule.Val = null;
      this.$http.post("http://localhost:8085/api/v1/rules", sending).then()
    },
    handleResize() {
      this.browserWidth = window.innerWidth
      this.browserHeight = window.innerHeight
    },
    onGridReady(params) {
      this.gridApi = params.api;
    },
    onSelectionChanged() {
      const selectedRows = this.gridApi.getSelectedRows();
      if (selectedRows[0] !== undefined) {
        this.selectedTracID = selectedRows[0].ID;
        this.selectedTracName = selectedRows[0].Name;
        console.log("Starting connection to WebSocket Server")
        if (this.connection != null) {
          this.connection.close()
        }
        this.connection = new WebSocket("ws://localhost:8085/api/v1/info/"+this.selectedTracID)
        this.connection.onmessage = (event) => {
          this.setInfo(event.data)
        }
       }
    },
    downloadJournal() {
      this.$http.get("http://localhost:8085/api/v1/notes", {params: {
        tractorID: this.selectedTracID
        }}).then(
          // eslint-disable-next-line no-unused-vars
          response=> {
            const notes = response && response.data ? response.data : {}
            const a = document.createElement("a");
            const data = JSON.stringify(notes, null, '\t')
            const file = new Blob([data], {type: 'text/html'});
            a.href = URL.createObjectURL(file);
            a.download = `Журнал нарушений для ${this.wsData.Name}.txt`;
            a.click();
            // eslint-disable-next-line no-unused-vars
          })
    },
  }
}
</script>

<style scoped>
.allContainer {
  margin-left: 15px;
  display: flex;
  flex-direction: row;
}
.title {
  margin-bottom: 15px;
  font-size: 24px;
  font-weight: bold;
  white-space: nowrap;

}

.tracInfo {
  margin-left: 50px;

}
.upper-r-table {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.upper-r-table-button {
  font-size: 12px;
  margin-left: 20px;
  margin-bottom: 16px;
  box-shadow: 5px 5px 10px 0 #2c3e50;
}
.upper-r-table-button:hover {
  box-shadow: 5px 5px 10px 0 #ce2a2e;
}
.infoField {
  width: 100%;
  border-bottom: #ce2a2e 1px solid;
  margin-bottom: 20px;
}
.inputField {
  margin-bottom: 5px;
}
.ag-theme-balham {
  --ag-header-background-color: #ce2a2e;
  --ag-header-foreground-color: #ffffff;
  height: calc(80vh - 70px)
}
</style>