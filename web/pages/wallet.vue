<template>
  <div class="wrap__page">
    <Header />
    <div class="account__wrapper pt-2">
      <div class="container container-1300">
        <div class="row gx-3 gx-xl-4">
          <SideBar />
          <div class="col-md-9 col--right d-flex">
            <div class="account__content w-100">
              <div class="row mb-3 mb-md-5">
                <div class="col-lg-4">
                  <div class="el__title">In-game balance</div>
                  <div
                    class="account__box__item d-flex justify-content-center flex-column"
                  >
                    <p class="el__box__sub">ASC</p>
                    <p class="el__box__title">152.245.521</p>
                    <button class="btn btn-change">Claim all</button>
                  </div>
                </div>
                <div class="col-lg-4">
                  <div class="el__title">Wallet balance</div>
                  <div
                    class="account__box__item -bnb d-flex justify-content-center flex-column"
                  >
                    <p class="el__box__sub">ASC</p>
                    <p class="el__box__title">152.245.521</p>
                  </div>
                </div>
                <div class="col-lg-4 d-flex">
                  <div
                    class="account__box__item mt-auto d-flex justify-content-center flex-column"
                  >
                    <p class="el__box__sub">BUSD</p>
                    <p class="el__box__title">152.245.521</p>
                  </div>
                </div>
              </div>

              <div class="mb-4">
                <div class="el__title mb-2">In-game balance</div>
                <p class="el__sub">Total: 300 heroes</p>
              </div>

              <div class="account__table mb-3">
                <table class="table mb-4">
                  <thead>
                    <tr>
                      <th scope="col" class="td__id">Id</th>
                      <th scope="col" class="td__name">Heroes</th>
                      <th class="d-lg-none" />
                      <th scope="col" class="td__claimable">Claimable</th>
                      <th scope="col" class="td__tax">
                        Tax
                        <i
                          class="icon-info-circle"
                          data-bs-toggle="tooltip"
                          data-bs-placement="top"
                          title="Initial tax will be 50% of the transaction value and gradually reduce 10% in a daily manner"
                        />
                      </th>
                      <th scope="col" class="td__stt" />
                    </tr>
                  </thead>

                  <tbody>
                    <tr v-for="i in 10" :key="i" class="js-tr-toggle">
                      <td scope="row" class="td__id">#6325</td>
                      <td class="td__name">
                        <div class="d-flex align-items-center">
                          <div class="td__thumb">
                            <img
                              src="/assets/images/market/table-user-avatar.png"
                              alt=""
                            />
                          </div>
                          <p>Magic force</p>
                        </div>
                      </td>
                      <td class="td__claimable">12,545,121 ASC</td>
                      <td class="td-toggle d-lg-none">
                        <i class="ic-arrow-down" />
                      </td>
                      <td class="td__tax">50%</td>
                      <td class="td__stt -claimed">
                        <button :class="['el--claim', { '-pending': i === 1 }]">
                          {{ i === 1 ? "Pending" : "Claim" }}
                        </button>
                      </td>
                    </tr>
                  </tbody>
                </table>

                <nav
                  class="navigation paging-navigation d-lg-flex align-items-center"
                  role="navigation"
                >
                  <div class="el__label text-center text-lg-start mb-3 mb-lg-0">
                    Showing 1 to 25 of 100 entries
                  </div>
                  <div
                    class="pagination loop-pagination ms-auto justify-content-center justify-content-lg-start"
                  >
                    <a class="prev page-numbers" href=""
                      ><i class="icon-arrow-prev"
                    /></a>
                    <span aria-current="page" class="page-numbers current"
                      >1</span
                    >
                    <a class="page-numbers" href="">2</a>
                    <a class="page-numbers" href="">3</a>
                    <a class="page-numbers" href="">4</a>
                    <a class="next page-numbers" href=""
                      ><i class="icon-arrow-next"
                    /></a>
                  </div>
                  <!-- .pagination -->
                </nav>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import $ from "jquery";
import SideBar from "@/components/account/side-bar.vue";
import Header from "@/components/account/header.vue";

export default {
  layout: "market",
  components: {
    SideBar,
    Header,
  },
  mounted() {
    if (this.$device.isMobile) {
      $(".js-tr-toggle").click(function (e) {
        $(this).next().toggle();
      });
    }
  },
  methods: {},
};
</script>
