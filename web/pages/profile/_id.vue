<template>
  <div>
    <div class="wrap__page -market">
      <Header />
      <div class="market__content pt-4">
        <div class="nft__single__back">
          <div class="container container-1200">
            <n-link to="/profile">Back</n-link>
          </div>
        </div>
        <div class="container container-800">
          <div class="row gx-4">
            <div class="col-md-36">
              <div class="nft__single__thumb">
                <div class="dnfix__thumb -contain">
                  <img
                    src="/assets/images/market/market-wizard-01.jpg"
                    alt=""
                  >
                </div>
              </div>
            </div>
            <div class="col-md-64">
              <div class="nft__single__summary">
                <div class="d-flex mb-4 pb-1">
                  <div class="">
                    <h1 class="el__title">Wizard</h1>
                    <div class="el__id">#4515</div>
                  </div>
                  <div class="ms-auto">
                    <div class="el__label">Rare</div>
                  </div>
                </div>

                <div class="mb-3">
                  <button
                    class="btn btn-primary w-100 -medium w-100"
                    data-bs-toggle="modal"
                    data-bs-target="#purchaseModal"
                  >
                    Sell now
                  </button>
                </div>

                <div class="d-flex gap-3 mb-4 pb-2">
                  <button
                    class="btn btn-secondary w-100 -medium w-100 mw-0"
                    data-bs-toggle="modal"
                    data-bs-target="#rentModal"
                  >
                    Make a rent
                  </button>
                  <button
                    class="btn btn-secondary w-100 -medium w-100 mw-0"
                    data-bs-toggle="modal"
                    data-bs-target="#purchaseModal"
                  >
                    Go on loan
                  </button>
                </div>

                <div class="single__tabs mb-4">
                  <div class="nav__wrap">
                    <ul
                      id="nftsingleTab"
                      class="nav flex-nowrap text-nowrap"
                      role="tablist"
                    >
                      <li class="nav-item" role="presentation">
                        <button
                          id="nftsingle1-tab"
                          class="nav-link active"
                          data-bs-toggle="tab"
                          data-bs-target="#nftsingle1"
                          type="button"
                          role="tab"
                          aria-controls="nftsingle1"
                          aria-selected="true"
                        >
                          Information
                        </button>
                      </li>
                      <li class="nav-item" role="presentation">
                        <button
                          id="nftsingle2-tab"
                          class="nav-link"
                          data-bs-toggle="tab"
                          data-bs-target="#nftsingle2"
                          type="button"
                          role="tab"
                          aria-controls="nftsingle2"
                          aria-selected="false"
                        >
                          Earning History
                        </button>
                      </li>
                      <li class="nav-item" role="presentation">
                        <button
                          id="nftsingle3-tab"
                          class="nav-link"
                          data-bs-toggle="tab"
                          data-bs-target="#nftsingle3"
                          type="button"
                          role="tab"
                          aria-controls="nftsingle3"
                          aria-selected="false"
                        >
                          Trading History
                        </button>
                      </li>
                    </ul>
                  </div>
                  <div id="nftsingleTabContent" class="tab-content">
                    <div
                      id="nftsingle1"
                      class="tab-pane fade show active"
                      role="tabpanel"
                      aria-labelledby="nftsingle1-tab"
                    >
                      <div class="table-responsive mb-4">
                        <table class="table">
                          <tr>
                            <td>Owner</td>
                            <td>0xCB4f...b2fDD1</td>
                          </tr>
                          <tr>
                            <td>Class</td>
                            <td>Wizard</td>
                          </tr>
                          <tr>
                            <td>Rarity</td>
                            <td>Common</td>
                          </tr>
                          <tr>
                            <td>Awaken level</td>
                            <td>II</td>
                          </tr>
                          <tr>
                            <td>Level</td>
                            <td>24</td>
                          </tr>
                        </table>
                      </div>

                      <div class="el__items">
                        <div class="el__items__title font18 mb-3">
                          Attributes
                        </div>

                        <div class="row mb-3">
                          <div class="col-12">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Overview</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-hourglass me-2" /> 1200 points
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>

                        <div class="row mb-3">
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Damage</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-sword me-2" /> 215
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Stamina</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-heart me-2" /> 4100
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Mana</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-mana me-2" /> 420
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Attack speed</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-star me-2" /> 2.5
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Critical</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-star me-2" /> 45%
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">DPS</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-star me-2" /> 45%
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>

                        <div class="row">
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Aura</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-key me-2" /> ICE
                                </div>
                              </div>
                            </div>
                          </div>
                          <div class="col-4">
                            <div class="el__item mb-3">
                              <div class="el__item__label">Ultimate</div>
                              <div class="el__item__value">
                                <div class="d-flex align-items-center">
                                  <i class="ic-cup me-2" /> Hunter’s Fury
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div
                      id="nftsingle2"
                      class="tab-pane fade"
                      role="tabpanel"
                      aria-labelledby="nftsingle2-tab"
                    >
                      <div class="table-responsive">
                        <table class="table">
                          <tr>
                            <td>Price</td>
                            <td>0xCB4Db2D1...</td>
                          </tr>
                          <tr>
                            <td>ROI</td>
                            <td>120%</td>
                          </tr>
                          <tr>
                            <td>Rarity</td>
                            <td>Common</td>
                          </tr>
                          <tr>
                            <td>Awaken level</td>
                            <td>II</td>
                          </tr>
                          <tr>
                            <td>Level</td>
                            <td>24</td>
                          </tr>
                        </table>
                      </div>
                    </div>
                    <div
                      id="nftsingle3"
                      class="tab-pane fade"
                      role="tabpanel"
                      aria-labelledby="nftsingle3-tab"
                    >
                      <div class="table-responsive">
                        <table class="table text-nowrap">
                          <thead>
                            <tr>
                              <th>Time</th>
                              <th>From</th>
                              <th>To</th>
                            </tr>
                          </thead>
                          <tbody>
                            <tr>
                              <td>hh:mm - dd/mm/yyyy</td>
                              <td>0xCB4f...b2fDD1</td>
                              <td>0xCB4f...b2fDD1</td>
                            </tr>
                            <tr>
                              <td>hh:mm - dd/mm/yyyy</td>
                              <td>0xCB4f...b2fDD1</td>
                              <td>0xCB4f...b2fDD1</td>
                            </tr>
                            <tr>
                              <td>hh:mm - dd/mm/yyyy</td>
                              <td>0xCB4f...b2fDD1</td>
                              <td>0xCB4f...b2fDD1</td>
                            </tr>
                          </tbody>
                        </table>
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
    <!-- .wrap__page -->

    <!-- Modal -->
    <div
      id="purchaseModal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="purchaseModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="purchaseModalLabel" class="modal-title">Unlist</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You are about to purchase</p>
              <p class="text-white">Wizard - #535353535</p>
            </div>
            <div class="d-flex mb-2">
              <div>Your balance:</div>
              <div class="ms-auto text-white">5215 ASC</div>
            </div>
            <div class="d-flex mb-4">
              <div>Fee:</div>
              <div class="ms-auto text-white">0,02414 ASC</div>
            </div>

            <div class="form-check mb-4">
              <input
                id="flexCheckDefault2"
                class="form-check-input"
                type="checkbox"
                value=""
              >
              <label class="form-check-label" for="flexCheckDefault2">
                I understand that after successfully bought, Hero level will be
                reset to nearest awaken level.
              </label>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn btn-secondary w-100 -medium mw-0"
              data-bs-dismiss="modal"
            >
              Cancel
            </button>
            <button
              type="button"
              class="btn btn-primary w-100 -medium mw-0"
              data-bs-toggle="modal"
              data-bs-target="#purchase2Modal"
            >
              Unlist now
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div
      id="purchaseModal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="purchaseModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="purchaseModalLabel" class="modal-title">Sell hero</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You are about to list this for sale Wizard #535353535</p>
            </div>
            <div class="el__input d-flex align-items-center mb-4">
              <input
                type="text"
                class="form-control"
                placeholder="Type your price"
              >
              <div class="el__input--prefix">
                <i class="ic-token me-2" />ASC
              </div>
            </div>
            <div class="d-flex mb-2">
              <div>Your balance:</div>
              <div class="ms-auto text-white">5215 ASC</div>
            </div>
            <div class="d-flex mb-4">
              <div>Service fee:</div>
              <div class="ms-auto text-white">0,02414 ASC</div>
            </div>

            <div class="form-check mb-4">
              <input
                id="flexCheckDefault"
                class="form-check-input"
                type="checkbox"
                value=""
              >
              <label class="form-check-label" for="flexCheckDefault">
                I understand that after successfully bought, Hero level will be
                reset to nearest awaken level.
              </label>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn btn-secondary w-100 mw-0 -medium"
              data-bs-dismiss="modal"
            >
              Cancel
            </button>
            <button
              type="button"
              class="btn btn-primary w-100 mw-0 -medium"
              data-bs-toggle="modal"
              data-bs-target="#purchase2Modal"
            >
              Sell now
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal SEll Step2 -->
    <div
      id="purchase2Modal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="purchase2ModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="purchase2ModalLabel" class="modal-title">Complete</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You are about to list this for sale Wizard #535353535</p>
            </div>

            <div class="d-flex mb-2">
              <div>Battle for rent:</div>
              <div class="ms-auto text-white">1,000 matches</div>
            </div>
            <div class="d-flex mb-2">
              <div>Price:</div>
              <div class="ms-auto text-white">5,215 ASC</div>
            </div>
            <div class="d-flex mb-4">
              <div>Maxium day to keep:</div>
              <div class="ms-auto text-white">256 days</div>
            </div>
            <div class="d-flex mb-4">
              <div>Fee:</div>
              <div class="ms-auto text-white">0,02414 ASC</div>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn btn-primary w-100 -medium"
              data-bs-toggle="modal"
              data-bs-target="#purchase2Modal"
            >
              OK
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div
      id="rentModal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="rentModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="rentModalLabel" class="modal-title">Lease</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You are about to list this for lease Wizard #535353535</p>
            </div>
            <div class="el__input d-flex align-items-center mb-4">
              <input
                type="text"
                class="form-control"
                placeholder="Type your price per match"
              >
              <div class="el__input--prefix">
                <i class="ic-token me-2" />ASC
              </div>
            </div>

            <div class="el__input -full d-flex align-items-center mb-4">
              <input
                type="text"
                class="form-control"
                placeholder="Match for lease"
              >
            </div>

            <div class="d-flex mb-2">
              <div>Your balance:</div>
              <div class="ms-auto text-white">5215 ASC</div>
            </div>
            <div class="d-flex mb-4">
              <div>Maxium day to keep:</div>
              <div class="ms-auto text-white">12 days</div>
            </div>

            <div class="alert alert-danger mb-4" role="alert">
              <strong>Oops ! Something went wrong</strong>
              <p>Minimum renting price: 40</p>
            </div>

            <div class="form-check mb-4">
              <input
                id="flexCheckDefault2"
                class="form-check-input"
                type="checkbox"
                value=""
              >
              <label class="form-check-label" for="flexCheckDefault2">
                I understand that after successfully bought, Hero level will be
                reset to nearest awaken level.
              </label>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn btn-secondary w-100 mw-0 -medium"
              data-bs-dismiss="modal"
            >
              Cancel
            </button>
            <button
              type="button"
              class="btn btn-primary w-100 mw-0 -medium"
              data-bs-toggle="modal"
              data-bs-target="#rent2Modal"
            >
              Sell now
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Step 2-->
    <div
      id="rent2Modal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="rent2ModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="rent2ModalLabel" class="modal-title">Lease</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You are about to lease</p>
              <p class="text-white">Wizard - #535353535</p>
            </div>

            <div class="el__input -full d-flex align-items-center mb-4">
              <input
                type="text"
                class="form-control"
                placeholder="Enter a number of match"
              >
            </div>

            <div class="d-flex mb-2">
              <div>Price per match:</div>
              <div class="ms-auto text-white">5,215 ASC</div>
            </div>
            <div class="d-flex mb-2">
              <div>Day to keep:</div>
              <div class="ms-auto text-white">n/a</div>
            </div>
            <div class="d-flex mb-2">
              <div>Total:</div>
              <div class="ms-auto text-white">n/a</div>
            </div>
            <hr>

            <div class="d-flex mb-2">
              <div>Your balance:</div>
              <div class="ms-auto text-white">5215 ASC</div>
            </div>
            <div class="d-flex mb-4">
              <div>Service fee:</div>
              <div class="ms-auto text-white">0,02414 ASC</div>
            </div>

            <div class="form-check mb-4">
              <input
                id="flexCheckDefault1"
                class="form-check-input"
                type="checkbox"
                value=""
              >
              <label class="form-check-label" for="flexCheckDefault1">
                I understand that after successfully bought, Hero level will be
                reset to nearest awaken level.
              </label>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn btn-secondary w-100 mw-0 -medium"
              data-bs-dismiss="modal"
            >
              Cancel
            </button>
            <button
              type="button"
              class="btn btn-primary w-100 mw-0 -medium"
              data-bs-toggle="modal"
              data-bs-target="#rent3Modal"
            >
              Check out
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Step 3-->
    <div
      id="rent3Modal"
      class="modal -buy fade"
      tabindex="-1"
      aria-labelledby="rent3ModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 id="rent3ModalLabel" class="modal-title">Lease</h5>
          </div>
          <div class="modal-body">
            <div class="el__text mb-4">
              <p>You have successfully listed for lease</p>
              <p class="text-white">Wizard - #535353535</p>
            </div>

            <div class="d-flex mb-2">
              <div>Matches:</div>
              <div class="ms-auto text-white">21</div>
            </div>
            <div class="d-flex mb-2">
              <div>Price per match:</div>
              <div class="ms-auto text-white">5,215 ASC</div>
            </div>
            <div class="d-flex mb-2">
              <div>Day to keep:</div>
              <div class="ms-auto text-white">n/a</div>
            </div>
            <div class="d-flex mb-2">
              <div>Total:</div>
              <div class="ms-auto text-white">n/a</div>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button type="button" class="btn btn-primary w-100 mw-0 -medium">
              OK
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '~/components/market/header.vue'

export default {
  layout: 'market',
  components: {
    Header
  },
  mounted () {},
  methods: {}
}
</script>
