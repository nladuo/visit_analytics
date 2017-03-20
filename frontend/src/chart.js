export const drawChart = (dates, counts) => {
  if ($(window).width() < 768) {
    $("#record_chart").css("width", ($("#detailModal").width() - 60) + "px")
  }

  var chart = echarts.init(document.getElementById('record_chart'));
  var option = {
      tooltip : {
          trigger: 'axis'
      },
      legend: {
          data:['网站访问量']
      },
      grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
      },
      xAxis : [{
          type : 'category',
          data : dates,
          splitLine: {
              show: false
          },
      }],
      yAxis : [{
          type : 'value',
          name : '访问量',

      }],
      series : [{
          name:'访问量',
          type:'line',
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
  chart.setOption(option);
  window.onresize = () => {
    if ($(window).width() < 768) {
      $("#record_chart").css("width", ($("#detailModal").width() - 60) + "px")
    } else {
      $("#record_chart").css("width", "560px")
    }
    chart.resize();
  }
}
