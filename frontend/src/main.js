"use strict"

import Api from "./utils/api"

Vue.config.delimiters = ['[[', ']]']

new Vue({
  el: '#app',
  data: {
    hosts: [],
    pages: [],
    selectedHost: ""
  },
  ready() {
    this.getHosts()

  },
  methods: {
    getHosts() {
      Api.get("/api/hosts", {}, (data) => {
        if (data != null) {
          this.hosts = data.data;
        }
      })

    },

    getPages() {
      console.log(this.selectedHost);
      Api.get("/api/pages", {host: this.selectedHost}, (data) => {
        if (data != null) {
          this.pages = data.data;
        }
      })

    },

    getDailyRecords() {

    }
  }
})
