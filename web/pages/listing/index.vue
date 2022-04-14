<template>
  <div class="wrap__page">
    <Header />
    <div class="account__wrapper pt-2">
      <div class="container container-1300">
        <div class="row gx-3 gx-xl-4">
          <SideBar />
          <div class="col-md-9 col--right d-flex">
            <div class="account__content w-100">
              <div class="market__content">
                <div class="market__content__top d-flex align-items-center">
                  <p class="el__order__total d-none d-sm-block">
                    Listing on marketplace
                  </p>

                  <div class="ms-auto">
                    <div class="d-flex align-items-center">
                      <span class="el__order__label me-3 d-none d-xl-block"
                        >Sort by:</span
                      >
                      <div class="btn-group">
                        <button
                          id="dropdownMenuSortby"
                          class="btn__sortby dropdown-toggle"
                          type="button"
                          data-bs-toggle="dropdown"
                          data-bs-auto-close="outside"
                          aria-expanded="false"
                        >
                          {{ getSortSelected() }}
                        </button>
                        <ul
                          class="dropdown-menu dropdown-menu-end dropdown-menu-dark"
                          aria-labelledby="dropdownMenuSortby"
                          style=""
                        >
                          <li v-for="(item, index) in sortOptions" :key="index">
                            <a
                              class="dropdown-item"
                              href="#"
                              @click="sort = item.value"
                              >{{ item.label }}</a
                            >
                          </li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row gx-3 gx-sm-4 mb-4 justify-content-center">
                  <div v-for="i in 20" :key="i" class="el__col col">
                    <div class="market__item ef--shine">
                      <n-link :to="`/listing/${i}`">
                        <div class="market__item__thumb">
                          <div class="dnfix__thumb -contain">
                            <img
                              src="/assets/images/market/market-wizard-01.jpg"
                              alt=""
                            />
                          </div>
                          <div class="market__item__label">45125</div>
                        </div>
                        <div class="market__item__meta">
                          <div class="d-flex mb-2">
                            <h3 class="market__item__title text__truncate">
                              Wizard
                            </h3>
                            <p class="market__item__id ms-auto">#ID</p>
                          </div>

                          <div class="market__item__price d-flex">
                            <div class="-label">Price</div>
                            <div class="-price ms-auto text-end">
                              <p class="-price-bnb">3.24 BNB</p>
                              <p class="-price-usd">$1,200 USD</p>
                            </div>
                          </div>
                        </div>
                      </n-link>
                    </div>
                  </div>
                </div>

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
                    <span class="page-numbers dots">…</span>
                    <a class="page-numbers" href="">7</a>
                    <a class="page-numbers" href="">8</a>
                    <a class="page-numbers" href="">9</a>
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
import SideBar from "~/components/account/side-bar.vue";
import Header from "~/components/account/header.vue";

export default {
  layout: "market",
  components: {
    SideBar,
    Header,
  },
  data() {
    return {
      sort: 0,
      sortOptions: [
        { label: "Lowest to highest", value: 0 },
        { label: "Lowest ID", value: 1 },
        { label: "Highest ID", value: 2 },
        { label: "Lowest Price", value: 3 },
        { label: "Highest Price", value: 4 },
        { label: "Latest", value: 5 },
      ],
    };
  },
  watch: {
    sort(val) {
      // handle sort at here
      console.log(val);
    },
  },
  mounted() {},
  methods: {
    getSortSelected() {
      const sort = this.sortOptions.find((item) => item.value === this.sort);
      return sort.label;
    },
  },
};
</script>
