<!-- class 会引用下面格式定义形式及尺寸 -->
<template>
  <a-layout class="container">
    <!-- 相当于直接引入 Nav.vue 文件 -->
    <Nav></Nav>

    <a-layout>

      <a-layout-header class="headerBtn">
        <Header></Header>>
      </a-layout-header>

      <a-layout-content>Content</a-layout-content>

      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>

    </a-layout>
  </a-layout>
</template>

<script>
  import Nav from '@/components/admin/Nav'
  import Footer from '@/components/admin/Footer'
  import Header from '@/components/admin/Header'
  export default {
    components: { Nav, Footer, Header },
  }
</script>

<style scoped>
/* body,html{
      width: 100%;
      height: 100%;
    } */
.container{
  
  /* width: 120px;
  height: 200px; */
  /* width: max-content;
  height: max-content; */
  /* position:absolute; */
  width: 100%;
  height: 100%;
      /* background-color:burlywood; */
  /* display: flex; */
      /* justify-content: center;
      align-items: center; */ 
}

.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

</style>
