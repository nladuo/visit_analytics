export default {
  get(url, data, _emit) {
    this.request("GET", url, data, _emit);
  },
  post(url, data, _emit) {
    this.request("POST", url, data, _emit);
  },

  request(method, url, data, _emit) {
    url = "/api" + url

    $('#loading').css('display','block');

    $.ajax({
      type : method,
      url : url,
      dataType: 'json',
      data : data,
      async: true,
      success(data) {
        _emit(data);
        $('#loading').css('display', 'none');
      },
      error() {
        _emit(null);
        $('#loading').css('display', 'none');
      }
    });
  }
}
