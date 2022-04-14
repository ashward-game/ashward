<template>
  <div class="wrap__page">
    <Header />
    <div class="account__wrapper pt-2">
      <div class="container container-1300">
        <div class="row gx-3 gx-xl-4">
          <SideBar />
          <div class="col-md-9 col--right d-flex">
            <div class="account__content w-100">
              <div class="row">
                <div class="col-lg-8 mb-4 mb-lg-0">
                  <div class="el__title mb-2">Basic Activities</div>
                  <p class="el__sub">
                    Below shows your attempted trasactions but they could fail
                    or be pending See the real status on Block Explore. Relea
                    see this arbcle for more info.
                  </p>
                </div>
                <div class="col-lg-4 text-end text-sm-start text-lg-end">
                  <div class="btn-group">
                    <button
                      id="dropdownMenuSortby"
                      class="btn__sortby dropdown-toggle"
                      type="button"
                      data-bs-toggle="dropdown"
                      data-bs-auto-close="outside"
                      aria-expanded="false"
                    >
                      {{ getFilterSelected() }}
                    </button>
                    <ul
                      class="dropdown-menu dropdown-menu-end dropdown-menu-dark"
                      aria-labelledby="dropdownMenuSortby"
                      style=""
                    >
                      <li
                        v-for="item in filterOptions"
                        :key="item.value"
                        @click="filter = item.value"
                      >
                        <a class="dropdown-item" href="#">{{ item.label }}</a>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
              <hr class="my-4" />
              <div v-if="filter !== 0" class="account__box">
                <div class="activity__box--none">
                  You don’t have any transaction
                </div>
              </div>
              <div v-else class="account__box">
                <div class="activity__box">
                  <h3 class="activity__box__title">20/12/2022</h3>
                  <div class="activity__box__item">
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You received
                          <span class="-color-yellow">+1300</span> ASC from
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You sent <span class="-color-red">-1576</span> ASC to
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img1.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You attempted to start a sale on
                          <span class="-color-yellow">Wizard #58659</span>
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img1.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You attempted to start a sale on
                          <span class="-color-yellow">Wizard #58659</span>
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img1.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You attempted to start a sale on
                          <span class="-color-yellow">Wizard #58659</span>
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img1.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You attempted to start a sale on
                          <span class="-color-yellow">Wizard #58659</span>
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You sent <span class="-color-red">-1576</span> ASC to
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You received
                          <span class="-color-yellow">+1300</span> ASC from
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="activity__box">
                  <h3 class="activity__box__title">19/12/2022</h3>
                  <div class="activity__box__item">
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img1.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You attempted to start a sale on
                          <span class="-color-yellow">Wizard #58659</span>
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You sent <span class="-color-red">-1576</span> ASC to
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="el__item">
                      <div class="el__item__date">hh:mm:ss</div>
                      <div class="el__item__thumb">
                        <img
                          src="/assets/images/market/activity-img2.png"
                          alt=""
                        />
                      </div>
                      <div class="el__item__meta">
                        <div class="el__item__title">
                          You received
                          <span class="-color-yellow">+1300</span> ASC from
                          0xaaaf...2cbcdc
                        </div>
                        <div class="el__item__sub">
                          <div class="d-none d-sm-block">
                            View transaction on BSC scan
                            <a href="" class="el__item__share"
                              ><i class="icon-copy"
                            /></a>
                          </div>
                          <div class="d-sm-none">
                            <a href="">View on BscScan</a>
                          </div>
                        </div>
                      </div>
                    </div>
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
      filter: 0,
      filterOptions: [
        { label: "Last 7 day", value: 0 },
        { label: "...", value: 1 },
        { label: "...", value: 2 },
      ],
    };
  },
  watch: {
    filter(val) {
      // handle sort at here
      console.log(val);
    },
  },
  mounted() {},
  methods: {
    getFilterSelected() {
      const filter = this.filterOptions.find(
        (item) => item.value === this.filter
      );
      return filter.label;
    },
  },
};
</script>
