<template>
<div class="allContainer">
  <!-- Левая панель с тракторами -->
  <div class="tracsList">
    <div class="title">
      Тракторы:
    </div>
    <div>
      <ag-grid-vue
          class="ag-theme-balham"
          :columnDefs="columns"
          :rowData="rows"
          :rowSelection="rowSelection"
          style="width: 252px"
          @selection-changed="onSelectionChanged()"
          @grid-ready="onGridReady"
      />
    </div>
  </div>

  <!-- Правая панель с информацией по трактору -->
  <div class="tracInfo" v-if=" selectedTracID !== '' ">
    <div class="upper-r-table">

      <div class="title" v-if="browserWidth > 1170">
        Информация о тракторе №{{selectedTracID}} - {{selectedTracName}}:
      </div>
      <div class="upper-r-table-btns">
        <b-button
            variant="danger" class="upper-r-table-button"
            @click="showAddRuleModal = true"
            style="margin-left: 0"
        >
          <b-icon-plus-circle/> Добавить правило
        </b-button>
        <b-button variant="danger" class="upper-r-table-button">
          <b-icon-journal-text/> Журнал нарушений
        </b-button>

      </div>

    </div>
    <ag-grid-vue
        class="ag-theme-balham"
        :columnDefs="columnsInfo"
        :rowData="rowsInfo"
        style=" width: 100%"

    />
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
</div>
</template>

<script>
import "ag-grid-community/styles/ag-grid.css";
import "ag-grid-community/styles/ag-theme-balham.css";
import {AgGridVue} from "ag-grid-vue";
export default {
  name: "AllView",
  components: {
    AgGridVue
  },
  data() {return {
    showAddRuleModal: false,
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
    rowSelection: 'single',
    selectedTracID: '',
    selectedTracName: '',
    columns: [
      {field: "id", headerName: 'Номер', sortable: true, width: 100},
      {field: "name", headerName: 'Название', sortable: true, width: 150}
    ],
    rows: [
      {id: 1, name: 'Трактор'},
      {id: 2, name: 'Кировец'},
      {id: 3, name: 'Череповец'},
      {id: 4, name: 'Череповец1'},
      {id: 5, name: 'Череповец2'},
      {id: 6, name: 'Череповец3'},
      {id: 7, name: 'Череповец4'},
      {id: 8, name: 'Череповец5'},
      {id: 9, name: 'Череповец6'},
      {id: 10, name: 'Череповец7'},
      {id: 11, name: 'Череповец8'},
      {id: 12, name: 'Череповец9'},
      {id: 13, name: 'Череповец10'},
      {id: 14, name: 'Череповец11'},
      {id: 15, name: 'Череповец12'},
      {id: 16, name: 'Череповец13'},
      {id: 17, name: 'Череповец14'},
      {id: 18, name: 'Череповец15'},
      {id: 19, name: 'Череповец16'},
      {id: 20, name: 'Череповец17'},
      {id: 21, name: 'Череповец18'},
      {id: 22, name: 'Череповец19'},
    ],
    selectField: [
      {text: "Выберите поле", value: null},
      {text: "SpeedRT", value: "SpeedRT"},
      {text: "EngineRPS", value: "EngineRPS"},
      {text: "FuelLevel", value: "FuelLevel"},
      {text: "FuelSpent", value: "FuelSpent"},
      {text: "FuelConsumption", value: "FuelConsumption"},
      {text: "EngineRegime", value: "EngineRegime"},
      {text: "OilTemperature", value: "OilTemperature"},
      {text: "EnvTemperature", value: "EnvTemperature"},
      {text: "RedLamp", value: "RedLamp"},
      {text: "WarnLamp", value: "WarnLamp"},
    ],
    columnsInfo: [
      {field: "SpeedRT"},
      {field: "EngineRPS"},
      {field: "FuelLevel"},
      {field: "FuelSpent"},
      {field: "FuelConsumption"},
      {field: "EngineRegime"},
      {field: "OilTemperature"},
      {field: "EnvTemperature"},
      {field: "RedLamp"},
      {field: "WarnLamp"},
    ],
    rowsInfo: [],
    browserWidth: 0,
    browserHeight: 0,
  }},
  created() {
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
  },
  onUnmounted() {
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
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
        ValInt: this.newRule.ValInt,
        ValFloat: this.newRule.ValFloat,
      }
      console.log("Sending: ", sending)
      this.showAddRuleModal = false
      this.newRule.Name = "";
      this.newRule.ValInt = 0;
      this.newRule.ValFloat = 0;
      this.newRule.FieldName = null;
      this.newRule.Val = null;
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
        this.selectedTracID = selectedRows[0].id;
        this.selectedTracName = selectedRows[0].name;
      }
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
  width: 75vw;
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
.inputField {
  margin-bottom: 5px;
}
.ag-theme-balham {
  --ag-header-background-color: #ce2a2e;
  --ag-header-foreground-color: #ffffff;
  height: calc(80vh - 70px)
}
</style>