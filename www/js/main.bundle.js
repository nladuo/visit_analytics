/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.l = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// identity function for calling harmony imports with the correct context
/******/ 	__webpack_require__.i = function(value) { return value; };

/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};

/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};

/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 3);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});
var drawChart = exports.drawChart = function drawChart(dates, counts) {
    var myChart = echarts.init(document.getElementById('record_chart'));
    var option = {
        tooltip: {
            trigger: 'axis'
        },
        legend: {
            data: ['网站访问量']
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: [{
            type: 'category',
            // boundaryGap : false,
            data: dates,
            splitLine: {
                show: false
            }
        }],
        yAxis: [{
            type: 'value',
            name: '访问量'

        }],
        series: [{
            name: '访问量',
            type: 'line',
            stack: '总量',
            symbolSize: 8,
            lineStyle: {
                normal: {
                    opacity: 1
                }
            },
            data: counts
        }]
    };
    myChart.setOption(option);
};

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var getToday = exports.getToday = function getToday() {
  var d = new Date();
  //fromat for yyyy-mm-dd
  var month = ((d.getMonth() + 1 + "").length == 1 ? "0" : "") + (d.getMonth() + 1);
  var date = ((d.getDate() + "").length == 1 ? "0" : "") + d.getDate();
  var date_str = d.getFullYear() + "-" + month + "-" + date;
  return date_str;
};

var initDatePicker = exports.initDatePicker = function initDatePicker() {
  $('#date_selected').datetimepicker({
    format: 'yyyy-mm-dd',
    weekStart: 1,
    autoclose: true,
    startView: 2,
    minView: 2,
    forceParse: false
  });

  $('#date_selected').datetimepicker('setEndDate', getToday());
};

/***/ }),
/* 2 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = {
  get: function get(url, data, _emit) {
    this.request("GET", url, data, _emit);
  },
  post: function post(url, data, _emit) {
    this.request("POST", url, data, _emit);
  },
  request: function request(method, url, data, _emit) {
    url = "/manage" + url;

    $('#loading').css('display', 'block');

    $.ajax({
      type: method,
      url: url,
      dataType: 'json',
      data: data,
      async: true,
      success: function success(data) {
        _emit(data);
        $('#loading').css('display', 'none');
      },
      error: function error() {
        _emit(null);
        $('#loading').css('display', 'none');
      }
    });
  }
};

/***/ }),
/* 3 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


var _api = __webpack_require__(2);

var _api2 = _interopRequireDefault(_api);

var _chart = __webpack_require__(0);

var _datepicker = __webpack_require__(1);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

Vue.config.delimiters = ['[[', ']]'];

new Vue({
  el: '#app',
  data: {
    hosts: [],
    pages: [],
    selectedUrl: "",
    selectedHost: "",
    selectedType: "0",
    dates: [],
    counts: []
  },
  ready: function ready() {
    (0, _datepicker.initDatePicker)();
    (0, _chart.drawChart)(this.dates, this.counts);
    this.getHosts();
  },

  methods: {
    getHosts: function getHosts() {
      var _this = this;

      _api2.default.get("/api/hosts", {}, function (data) {
        if (data != null) {
          _this.hosts = data.data;
        }
      });
    },
    getPages: function getPages() {
      var _this2 = this;

      console.log(this.selectedHost);
      _api2.default.get("/api/pages", { host: this.selectedHost }, function (data) {
        if (data != null) {
          _this2.pages = data.data;
        }
      });
    },
    getRecords: function getRecords() {
      var _this3 = this;

      _api2.default.get("/api/records", {
        date: $("#date_selected").val(),
        type: this.selectedType,
        url: this.selectedUrl
      }, function (data) {
        _this3.dates = [];
        _this3.counts = [];

        data.data.forEach(function (record) {
          _this3.dates.push(record.Date);
          _this3.counts.push(record.Count);
        });
        console.log(data);
        (0, _chart.drawChart)(_this3.dates, _this3.counts);
        $("#detailModal").modal('show');
      });
    },
    showDetail: function showDetail(url) {
      $("#date_selected").val((0, _datepicker.getToday)());
      this.selectedUrl = url;
      this.selectedType = "0";
      this.getRecords();
    }
  }
});

/***/ })
/******/ ]);