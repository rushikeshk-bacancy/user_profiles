<template>
  <div class="Background">
    <h3 class="Title">Add User Profile</h3>
    <form class="Title" @submit.prevent="uploadFileAndSave(file)">
      <div class="form-group">
        <label for="exampleFormControlInput1">First Name</label>
        <input
          v-model="user.firstName"
          type="text"
          maxlength="20"
          class="form-control"
          id="exampleFormControlInput1"
          placeholder="Enter Your First Name"
        />
      </div>
      <div class="form-group">
        <label for="exampleFormControlInput1">Last Name</label>
        <input
          v-model="user.lastName"
          type="text"
          maxlength="20"
          class="form-control"
          id="exampleFormControlInput1"
          placeholder="Enter Your Last Name"
        />
      </div>
      <div class="form-group">
        <label for="exampleFormControlInput1">Email Id</label>
        <input
          v-model="user.email"
          type="text"
          maxlength="20"
          class="form-control"
          id="exampleFormControlInput1"
          placeholder="Enter Your Email Id"
        />
      </div>
      <div class="form-group">
        <label for="exampleFormControlInput1">Country Name</label>
        <input
          v-model="user.country"
          type="text"
          maxlength="20"
          class="form-control"
          id="exampleFormControlInput1"
          placeholder="Enter Your Country Name"
        />
      </div>
      <div class="form-group">
        <label for="exampleFormControlFile1">Add User Profile Picture</label>
        <input
          @input="handleFileUpload()"
          ref="file"
          id="file"
          type="file"
          class="form-control-file"
        />
      </div>

      <div class="form-group">
        <button type="submit" class="btn btn-primary">Submit</button>
        <button
          type="reset"
          @click="reDirect()"
          class="btn btn-danger"
          style="margin-left: 10px"
        >
          Cancel
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  mounted() {},
  data() {
    return {
      user: {
        userId: "",
        firstName: "",
        lastName: "",
        country: "",
        dob: "",
        userPicture: "",
        gender: "M",
        email: "",
      },
      file: "",
    };
  },
  methods: {
    uploadFileAndSave(file) {
      if (file != "") {
        const formData = new FormData();
        formData.append("file", file);
        this.axios
          .post("http://localhost:4700/upload", formData, {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          })
          .then((resp) => {
            if (resp.status === 200) {
              console.log("File uploaded");
              this.user.userPicture = "http://localhost:4700/"+file.name
              this.saveData()
            }
          });
      }else{
        this.saveData()
      }
      
    },
    reDirect() {
      this.$router.push({ name: "Dashboard" });
    },
    saveData() {
      console.log(this.user)
      this.axios
        .post("http://localhost:4700/add", this.user)
        .then((response) => {
          if (response.request.status !== 200) {
            alert("Some Error Occured");
          } else {
            this.$router.push({ name: "Dashboard" });
          }
        });
    },
    handleFileUpload() {
      this.file = this.$refs.file.files[0];
    },
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
</style>
