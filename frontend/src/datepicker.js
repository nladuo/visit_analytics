export const getToday = () => {
  return (new Date()).Format("yyyy-MM-dd");
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
