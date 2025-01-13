
<template>
  <div style="margin-top: 20px">
    <a-descriptions :data="data" :column="2" bordered />
  </div>
</template>

<script>

export default {
  setup() {
    const data = [{
      label: '标题',
      value: 'Socrates',
    }, {
      label: '子工单号',
      value: '123-1234-1234',
    }, {
      label: '申请人',
      value: '123-1234-1234',
    }, {
      label: '创建时间',
      value: '123-1234-1234',
    }, {
      label: '审批人',
      value: '123-1234-1234',
    },{
      label: '更新时间',
      value: '123-1234-1234',
    },{
      label: '描述',
      value: 'Yingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, BeijingYingdu Building, Zhichun Road, Beijing'
    }];

    return {
      data,
    }
  },
}
</script>
