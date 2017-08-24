<template>
  <div class="documents">
      <form class="form-horizontal" @submit.prevent="submit">
         <div class="form-group">
            <label for="initiativeWarn" class="col-sm-2 control-label text-info">第一个文件的相关信息</label>
        </div> 
        <div class="clearfix"></div>
        <div class="form-group">
            <label for="initiativeRow" class="col-sm-2 control-label">起始行号</label>
            <div class="col-sm-10">
            <input type="text" class="form-control" name="initiativeRow" placeholder="1" value="2">
            </div>
        </div>
        <div class="form-group">
            <label for="initiativeColumn" class="col-sm-2 control-label">比较列号</label>
            <div class="col-sm-10">
            <input type="text" class="form-control" name="initiativeColumn" placeholder="1">
            </div>
        </div>
        <div class="form-group">
            <label for="PassiveWarn" class="col-sm-2 control-label text-info">第二个文件的相关信息</label>
        </div> 
        <div class="form-group">
            <label for="PassiveRow" class="col-sm-2 control-label">起始行号</label>
            <div class="col-sm-10">
            <input type="text" class="form-control" placeholder="1">
            </div>
        </div>
        <div class="form-group">
            <label for="PassiveColumn" class="col-sm-2 control-label">比较列号</label>
            <div class="col-sm-10">
            <input type="text" class="form-control" name="PassiveColumn" placeholder="1">
            </div>
        </div>
        <div class="form-group">
            <label for="SelectBlock" class="col-sm-2 control-label text-info">文件选择区域</label>
        </div>
        <div class="form-group">
            <label for="InitiativeFile" class="col-sm-2 control-label">选择文件一</label>
            <div class="col-sm-10">
            <input type="file" name="initiativeFile" placeholder="1">
            </div>
        </div>
        <div class="form-group">
            <label for="PassiveFile" class="col-sm-2 control-label">选择文件二</label>
            <div class="col-sm-10">
            <input type="file" name="passiveFile" placeholder="1">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-offset-2 col-sm-10">
            <button type="submit" class="btn btn-primary pull-right">提交</button>
            </div>
        </div>
      </form>
  </div>
</template>
<script>
export default {
  data () {
    return {
      input: ''
    }
  },
  methods: {
    submit: function (event) {
      console.log('====')
      this.$http.post('', event.target, function (data, staus, request) {}).error(function (data, status, request) {})
      console.log(this.$http)
    }
  }
}
</script>


