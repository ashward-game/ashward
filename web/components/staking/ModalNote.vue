<template>
  <!-- Modal noteModal -->
  <div
    id="noteModal"
    class="modal -stake fade"
    tabindex="-1"
    aria-hidden="true"
  >
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header justify-content-center">
          <h5 class="modal-title">Notes</h5>
        </div>
        <div class="modal-body">
          <div class="el__item">
            <p class="el__item__title">Gift box</p>
            <div class="el__item__content">
              <p>
                Gift box is an item rewarded to testnet event whitelist winners
                and anyone who participates in the ASC staking pool.
              </p>
              <p>
                Each gift box contains 1 NFT and can be redeemed immediately
                after joining staking. This NFT will be available in game for
                players to play testnet.
              </p>
              <p>
                For each minimum amount staked adhere specifically to every
                pool, user will receive a gift box.
              </p>
            </div>
          </div>

          <div class="el__item">
            <p class="el__item__title">Gift code</p>
            <div class="el__item__content">
              <p>
                Gift code is an item rewarded to who is eligible for gift box
                and extend to users who use reference code to invite more users
                to join the game.
              </p>
              <p>
                Gift code contains 1 item from the following pool: NFT, Free
                Awaken, Boxes of ASC, Pet (7 days), Decorative items.
              </p>
              <p>
                These gift codes can be exchanged to real-valued items when the
                mainnet version is released.
              </p>
            </div>
          </div>
        </div>
        <div class="modal-actions d-flex justify-content-center gap-3">
          <button type="button" class="btn-stake" data-bs-dismiss="modal">
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {}
</script>
