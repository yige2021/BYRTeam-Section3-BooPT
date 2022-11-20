<template>
  <div id="hpgBackground">
    <el-container id="hpgContainer">
      <el-header id="hpgHeader">
        <el-row :gutter="20">
          <el-col :span="20">
            <img src="../assets/minilogo.png"></img>
          </el-col>
          <el-col :span="4" id="rcol">
            <el-link id="quit" type="info" @click="quitEvent">注销</el-link>
          </el-col>
        </el-row>
      </el-header>
      <el-container>
        <el-aside id="hpgAside" width="200px">
          <el-menu>
            <el-menu-item index="1-1">我的书目</el-menu-item>
            <el-menu-item index="1-1">待开发</el-menu-item>
            <el-menu-item index="1-1">待开发</el-menu-item>
          </el-menu>
          容量：
          <el-progress id="storageVol" :percentage="50"></el-progress>
        </el-aside>
        <el-main id="hpgMain">
          <el-row id="updateRow" :gutter="20" >
              <el-upload
                class="upload-demo"
                action="https://jsonplaceholder.typicode.com/posts/"
                multiple>
                <el-button type="primary" icon="el-icon-upload">上传电子书</el-button>
              </el-upload>
          </el-row>
          <el-table
          id="bookList"
          :data="tableData"
          style="width: 100%">
            <el-table-column
              prop="title"
              label="标题"
              width="280">
            </el-table-column>
            <el-table-column
              prop="author"
              label="作者"
              width="300">
            </el-table-column>
            <el-table-column
              prop="publisher"
              label="出版社"
              width="180">
            </el-table-column>
            <el-table-column
              prop="date"
              label="修改日期">
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>


<script>
  export default {
    name: "Homepage.vue",
    data() {
      return {
        stuid: "!23",//
        fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}, {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}],
        tableData: [{
          title: '深入理解计算机系统',
          author: 'Randal E. Bryant & David R. O\'Hallaron',
          publisher: '机械工业出版社',
          date: '2022年11月18日',
        }]
      }
    },
    methods:{
      quitEvent() {
        console.log("sb")
        const self = this
        window.sessionStorage.clear();
        self.$router.push('/');
      }
    }
  }
</script>

<style>
  #hpgBackground {
    background-color: #FFFFFF;
    box-shadow:0 2px 12px 0 rgba(0, 0, 0, 0.1);
    width:100%;	
    height:100%;	
    position:fixed;
    background-size:100% 
  }
  #hpgContainer {
    height:100%;
  }
  #hpgHeader {
    box-shadow:0 2px 12px 0 rgba(0, 0, 0, 0.1);
    min-height: 70px;
  }
  #hpgAside {
    position: relative;
    height:100%;
    top: 15px;
    box-shadow:0 2px 12px 0 rgba(0, 0, 0, 0.1);
    text-align: center;
  }
  #hpgMain {
    background-color: #FFFFFF;
    text-align: left;
    position:relative;
    height: 100%;
    top:15px;
    left:10px;
  }
  #rcol {
    text-align: right;
  }
  #quit {
    top: 25px;
  }
  #storageVol {
    position: relative;
    width: 180px;
  }
  #updateRow {
    position: absolute;
    top: 1%;
    left: 3%;
  }
  #bookList {
    position: relative;
    top: 30px;
  }
</style>