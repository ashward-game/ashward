<template>
  <div class="wrap__page">
    <Header />
    <div class="account__wrapper pt-2">
      <div class="container container-1300">
        <div class="row gx-3 gx-xl-4">
          <SideBar />
          <div class="col-md-9 col--right d-flex">
            <div class="account__content w-100">
              <div class="el__title">Basic Information</div>

              <div class="el__box">
                <label class="form-label">Name</label>
                <div class="d-flex mb-4 gap-3">
                  <input
                    v-model="name"
                    type="text"
                    class="form-control input-style"
                    placeholder="Name"
                  />
                  <button class="btn btn-change" @click="changeName">
                    <span class="d-none d-sm-block"> Change name </span>
                    <span class="d-sm-none">Edit</span>
                  </button>
                </div>

                <div class="mb-4">
                  <label class="form-label">Email address</label>
                  <div class="el__email">
                    <input
                      v-model="email"
                      type="text"
                      class="form-control input-style -stt"
                      placeholder="Email address"
                    />
                    <div class="el__email--stt">
                      <i class="icon-check-circle" />
                      <span class="d-none d-sm-block">Veried</span>
                    </div>
                  </div>
                </div>
              </div>

              <hr class="my-5" />

              <div class="el__title">Account setting</div>

              <div class="row">
                <div class="col-lg-6 mb-4">
                  <div class="form-check form-switch">
                    <input
                      id="form-switch-noti1"
                      v-model="notification1"
                      class="form-check-input"
                      type="checkbox"
                    />
                    <label class="form-check-label" for="form-switch-noti1"
                      >Recived notification from Ashward</label
                    >
                  </div>
                </div>
                <div class="col-lg-6 mb-4">
                  <div class="form-check form-switch">
                    <input
                      id="form-switch-noti2"
                      v-model="notification2"
                      class="form-check-input"
                      type="checkbox"
                    />
                    <label class="form-check-label" for="form-switch-noti2"
                      >Recived notification from Ashward</label
                    >
                  </div>
                </div>
                <div class="col-lg-6 mb-4">
                  <div class="form-check form-switch">
                    <input
                      id="form-switch-noti3"
                      v-model="notification3"
                      class="form-check-input"
                      type="checkbox"
                    />
                    <label class="form-check-label" for="form-switch-noti3"
                      >Recived notification from Ashward</label
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SideBar from "@/components/account/side-bar.vue";
import Header from "@/components/account/header.vue";

export default {
  layout: "market",
  components: {
    SideBar,
    Header,
  },
  data() {
    return {
      name: "Sluthxie",
      email: "Sluthxie@teamdesign.sphoton.com",
      notification1: true,
      notification2: false,
      notification3: false,
    };
  },
  watch: {
    email(val) {
      // handle logic at here
      console.log(val);
    },
    notification1(val) {
      // handle logic at here
      console.log(val);
    },
    notification2(val) {
      // handle logic at here
      console.log(val);
    },
    notification3(val) {
      // handle logic at here
      console.log(val);
    },
  },
  methods: {
    changeName() {
      // handle logic at here
      console.log(this.name);
    },
  },
};
</script>
