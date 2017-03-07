export const getToday = () => {
  var d = new Date();
  //fromat for yyyy-mm-dd
  var month = (((d.getMonth() + 1) + "").length == 1 ? "0" : "") + (d.getMonth() + 1);
  var date = ((d.getDate() + "").length == 1 ? "0": "") + d.getDate();
  var date_str = d.getFullYear() + "-" + (month) + "-" + date;
  return date_str;
}

export const initDatePicker = () => {
  $('#date_selected').datetimepicker({
      format: 'yyyy-mm-dd',
      weekStart: 1,
      autoclose: true,
      startView: 2,
      minView: 2,
      forceParse: false
  });

  $('#date_selected').datetimepicker('setEndDate', getToday());
}
