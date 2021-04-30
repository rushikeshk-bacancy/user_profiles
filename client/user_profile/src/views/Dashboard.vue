<template>
  <div class="Background">
    <!-- <span class="mdi mdi-name"></span> -->
    <h3 class="Title">
      User Profiles
      <!-- <i class="mdi mdi-sort-alphabetical"></i> -->
      <span class="AddButton">
        <button class="btn btn-primary" @click="AddUser" style="margin-right: 10px">
          Add User
        </button>
      </span>
      <span class="Sort">
        <button v-on:click="ascending = !ascending" class="sort-button">
          <img
            v-if="ascending"
            src="@/assets/asc.jpg"
            style="width: 24px; height: 24px"
          />
          <img
            v-else
            src="@/assets/desc.jpg"
            style="width: 24px; height: 24px"
          />
        </button>
      </span>
      <span class="Search">
        <input placeholder="Search By First Name" v-model="searchText" type="text" />
      </span>
    </h3>
    <div class="card-bg">
      <div class="row" v-if="computeUsers.length">
        <div class="col-3" v-for="(user, index) in computeUsers" :key="index">
          <b-card
            tag="article"
            style="max-width: 20rem"
            class="mb-2"
          >
            <!-- :title="fullName(user)"
            :img-src="checkImage(user)" -->
          <img v-if="user.userPicture==''" src="../assets/default_avatar.jpg" style="width:250px" /><br/>
          <img v-if="user.userPicture!=''" :src="user.userPicture" style="width:250px" /><br/>
          <h5>{{fullName(user)}}</h5>
            <b-card-text>
             <span> Email : {{ user.email }} </span><br/>
             <span>Country : {{ user.country }}</span>
            </b-card-text>

            <button @click="EditUser(user.userId)" class="btn btn-primary" style="margin-left: 10px">
          Edit
        </button>
            <button @click="RemoveUser(user.userId)" class="btn btn-danger" style="margin-left: 10px">
          Delete
        </button>
          </b-card>
          <!-- <b-card
            v-else
            img-top
            tag="article"
            style="max-width: 20rem"
            class="mb-2"
          >
          <img src="../assets/default_avatar.jpg" style="width:250px" /><br/>
          <h5>{{fullName(user)}}</h5>

            <b-card-text>
             <span> Email  : {{ user.email }} </span><br/>
             <span>Country : {{ user.country }}</span>
            </b-card-text>

            <button @click="EditUser(user.userId)" class="btn btn-primary" style="margin-left: 10px">
          Edit
        </button>
            <button @click="RemoveUser(user.userId)" class="btn btn-danger" style="margin-left: 10px">
          Delete
        </button>
          </b-card> -->
        </div>
      </div>
      <div class="row" v-if="!computeUsers.length">
        <span>No data found</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Dashboard",
  components: {},
  data() {
    return {
      users: [],
      searchText: "",
      ascending: true,
    };
  },
  mounted() {
    this.axios.get("http://localhost:4700/view").then((response) => {
      this.users = response.data;
    });
  },
  computed: {
    computeUsers() {
      var tempUserArray = this.users;
      // Modified Code
      // Search
      if (this.searchText != "" && this.searchText) {
        tempUserArray = tempUserArray.filter((item) => {
          return item.firstName
            .toLowerCase()
            .includes(this.searchText.toLowerCase());
        });
      }

      // Sort by alphabetical order
      tempUserArray = tempUserArray.sort((a, b) => {
        let fa = a.firstName.toLowerCase(),
          fb = b.firstName.toLowerCase();

        if (fa < fb) {
          return -1;
        }
        if (fa > fb) {
          return 1;
        }
        return 0;
      });

      // Show sorted array in descending or ascending order
      if (!this.ascending) {
        tempUserArray.reverse();
      }
      return tempUserArray;
    },
  },
  methods: {
    fullName(user) {
      return user.firstName + " " + user.lastName;
    },
    RemoveUser(userId){
      this.axios.post("http://localhost:4700/delete",{ userId: userId }).then((response) => {
      this.users = response.data;
    });
    },
    EditUser(UserId){
      this.$router.push({
        name:"EditUser",params:{
          userId:UserId
        }
      })
    },
    AddUser(){
      this.$router.push({name:"AddUser"})
    },
    checkImage(user){
      if(user.userPicture!=""){
        return user.userPicture
      }else{
        return "../assets/default_avatar.jpg"
      }
    }
  },
};
</script>

<style>
.Background {
  height: 942px;
  left: 285px;
  right: 33px;
  top: 128px;
}
.Title {
  height: 24px;
  left: 32px;
  right: 690px;
  top: 32px;
  margin-top: 100px;
  margin-left: 100px;
  font-family: Mulish;
  font-style: normal;
  font-weight: bold;
  font-size: 19px;
  line-height: 24px;
  letter-spacing: 0.4px;
  color: #252733;
}
.Sort {
  width: 53px;
  height: 20px;
  left: 947px;
  top: 37px;
  font-weight: normal;
}
.Search {
  width: 53px;
  height: 20px;
  left: 947px;
  top: 37px;
  font-weight: normal;
  margin-left: 10px;
}
.card-bg {
  margin-top: 50px;
  margin-left: 100px;
}
.AddButton{
  margin-left: 1000px;
}
</style>
