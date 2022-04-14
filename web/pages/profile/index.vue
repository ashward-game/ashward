<template>
  <div class="wrap__page -market">
    <div class="market__mb d-md-none">
      <div class="container">
        <div class="d-flex align-items-center justify-content-between">
          <button
            class="btn-filter -filter dropdown-toggle"
            type="button"
            data-bs-toggle="dropdown"
            data-bs-auto-close="outside"
            aria-expanded="false"
          >
            {{ getSortTypeSelected() }}
          </button>
          <ul class="dropdown-menu dropdown-menu-start dropdown-menu-dark">
            <li v-for="(item, index) in sortTypeOptions" :key="index">
              <a
                class="dropdown-item"
                href="#"
                @click="sortType = item.value"
                >{{ item.label }}</a
              >
            </li>
          </ul>
          <div class="btn-filter js-widget--toggle">Filter</div>
        </div>
      </div>
    </div>

    <Header />

    <div class="market__content">
      <div class="container container-1200">
        <div class="row gx-3">
          <div class="col-12 col--left">
            <div class="market__sidebar">
              <div class="d-flex d-md-none pb-4">
                <div
                  class="widget-area__close ms-auto icon-close js-widget--toggle"
                />
              </div>
              <div class="market__sidebar__header">
                <div class="d-flex">
                  <div class="el__label">Filter</div>
                  <div class="el__reset ms-auto" @click="resetFilter">
                    Clear
                  </div>
                </div>
              </div>
              <div class="market__sidebar__content">
                <div class="market__widget -checkbox">
                  <div
                    class="market__widget__header -collapse d-flex align-items-center"
                  >
                    <a
                      class="market__widget__title"
                      data-bs-toggle="collapse"
                      href="#market-widget--collapse01"
                      role="button"
                      aria-expanded="true"
                      aria-controls="market-widget--collapse01"
                      >Class</a
                    >
                    <div
                      class="market__widget__reset ms-auto d-none d-lg-block"
                      @click="filter.class = []"
                    >
                      Reset
                    </div>
                  </div>

                  <div
                    id="market-widget--collapse01"
                    class="market__widget__content multi-collapse collapse show"
                  >
                    <ul class="js-checkbox-list">
                      <li v-for="item in classOptions" :key="item.value">
                        <label class="checkbox__item">
                          <input
                            v-model="filter.class"
                            type="checkbox"
                            :value="item.value"
                            class="js-checkall"
                          /><span class="checkmark" />
                          <span class="-label">{{ item.label }}</span>
                        </label>
                      </li>
                    </ul>
                  </div>
                </div>

                <div class="market__widget -checkbox">
                  <div
                    class="market__widget__header -collapse d-flex align-items-center"
                  >
                    <a
                      class="market__widget__title"
                      data-bs-toggle="collapse"
                      href="#market-widget--collapse01"
                      role="button"
                      aria-expanded="true"
                      aria-controls="market-widget--collapse01"
                      >Grade</a
                    >
                    <div
                      class="market__widget__reset ms-auto d-none d-lg-block"
                      @click="filter.grade = []"
                    >
                      Reset
                    </div>
                  </div>
                  <div
                    id="market-widget--collapse02"
                    class="market__widget__content multi-collapse collapse show"
                  >
                    <ul class="js-checkbox-list">
                      <li v-for="item in gradeOptions" :key="item.value">
                        <label class="checkbox__item">
                          <input
                            v-model="filter.grade"
                            type="checkbox"
                            :value="item.value"
                            class="js-checkall"
                          /><span class="checkmark" />
                          <span class="-label">{{ item.label }}</span>
                        </label>
                      </li>
                    </ul>
                  </div>
                </div>

                <div class="market__widget -level">
                  <div class="market__widget__header d-flex align-items-center">
                    <div class="market__widget__title">Level</div>
                    <div
                      class="market__widget__reset ms-auto"
                      @click="resetLevel"
                    >
                      Reset
                    </div>
                  </div>
                  <div class="market__widget__content">
                    <IonRangeSlider
                      v-model="rangeSlider"
                      :options="optionsRanges"
                      :value="'1;20'"
                    />
                  </div>
                </div>

                <div class="market__widget -strength">
                  <div class="market__widget__header d-flex align-items-center">
                    <div class="market__widget__title">Strength</div>
                    <div
                      class="market__widget__reset ms-auto"
                      @click="filter.strength = { min: 0, max: 0 }"
                    >
                      Reset
                    </div>
                  </div>
                  <div class="market__widget__content">
                    <div
                      class="d-flex align-items-center justify-content-between"
                    >
                      <input
                        v-model="filter.strength.min"
                        type="text"
                        placeholder="0000"
                      />
                      <div class="mx-2">-</div>
                      <input
                        v-model="filter.strength.max"
                        type="text"
                        placeholder="0000"
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-12 col--right">
            <div class="px-md-4">
              <div class="market__content__top d-flex align-items-center">
                <div class="d-none d-md-block">
                  <button
                    class="btn__sortby dropdown-toggle"
                    type="button"
                    data-bs-toggle="dropdown"
                    data-bs-auto-close="outside"
                    aria-expanded="false"
                  >
                    {{ getSortTypeSelected() }}
                  </button>
                  <ul
                    class="dropdown-menu dropdown-menu-start dropdown-menu-dark"
                  >
                    <li v-for="(item, index) in sortTypeOptions" :key="index">
                      <a
                        class="dropdown-item"
                        href="#"
                        @click="sortType = item.value"
                        >{{ item.label }}</a
                      >
                    </li>
                  </ul>
                </div>

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
                    <n-link :to="`/profile/${i}`">
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
</template>

<script>
import $ from "jquery";
import Header from "~/components/account/header.vue";
import IonRangeSlider from "~/components/market/IonRangeSlider.vue";

export default {
  layout: "market",
  components: {
    Header,
    IonRangeSlider,
  },
  data() {
    return {
      rangeSlider: 0,
      sort: 0,
      sortType: 0,
      sortOptions: [
        { label: "Lowest to highest", value: 0 },
        { label: "Lowest ID", value: 1 },
        { label: "Highest ID", value: 2 },
        { label: "Lowest Price", value: 3 },
        { label: "Highest Price", value: 4 },
        { label: "Latest", value: 5 },
      ],
      sortTypeOptions: [
        { label: "Heroes", value: 0 },
        { label: "Afon", value: 1 },
        { label: "Hefina", value: 2 },
        { label: "Idris", value: 3 },
        { label: "Macsen", value: 4 },
        { label: "...", value: 5 },
      ],
      filter: {
        class: [],
        grade: [],
        level: {
          min: 13,
          max: 38,
        },
        strength: {
          min: 0,
          max: 0,
        },
      },
      classOptions: [
        { label: "All class", value: 0 },
        { label: "Common", value: 1 },
        { label: "Rare", value: 2 },
        { label: "Epic", value: 3 },
        { label: "Legendary", value: 4 },
        { label: "Mythical", value: 5 },
      ],
      gradeOptions: [
        { label: "All class", value: 0 },
        { label: "Common", value: 1 },
        { label: "Rare", value: 2 },
        { label: "Epic", value: 3 },
        { label: "Legendary", value: 4 },
        { label: "Mythical", value: 5 },
      ],
      optionsRanges: {
        skin: "square",
        type: "double",
        min: 0,
        max: 50,
        from: 1,
        to: 20,
        grid: true,
      },
    };
  },
  watch: {
    sort(val) {
      // handle sort at here
      console.log(val);
    },
    sortType(val) {
      // handle sort at here
      console.log(val);
    },
    filter: {
      handler() {
        this.handleSubmit();
      },
      deep: true,
    },
    rangeSlider(val) {
      if (val) {
        const level = val.split(";");
        this.filter.level = { from: level[0], to: level[1] };
      }
    },
  },
  mounted() {
    $(".js-widget--toggle").on("click", function (e) {
      $(".market__sidebar").toggleClass("active");
      $("body").toggleClass("filter-open");
    });
  },
  methods: {
    getSortSelected() {
      const sort = this.sortOptions.find((item) => item.value === this.sort);
      return sort.label;
    },
    getSortTypeSelected() {
      const sortType = this.sortTypeOptions.find(
        (item) => item.value === this.sortType
      );
      return sortType.label;
    },
    resetFilter() {
      this.filter = {
        name: "",
        class: [],
        grade: [],
        level: {
          min: 13,
          max: 38,
        },
        strength: {
          min: 0,
          max: 0,
        },
      };
      this.resetLevel();
    },
    resetLevel() {
      this.filter.label = { from: 1, to: 20 };
      this.optionsRanges = {
        skin: "square",
        type: "double",
        min: 0,
        max: 50,
        from: 1,
        to: 20,
        grid: true,
      };
    },
    handleSubmit() {
      console.log(this.filter);
    },
  },
};
</script>
