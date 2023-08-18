<template>
  <v-row>

    <v-col cols="6">
      <v-text-field
          clearable
          label="Search stock symbol"
          v-model="symbol"
          @keyup.enter="search"
      />
    </v-col>

    <v-col cols="6">
      <v-btn
          variant="tonal"
          href="https://stockanalysis.com/stocks/"
          target="_blank"
          class="ma-4"
      >
        Stock symbols list
      </v-btn>
    </v-col>

  </v-row>


  <v-container class="wallboard" fluid>

    <v-row>

      <v-col
          v-for="(data, index) in companyData"
          :key="data.title"
          :cols="getCols(data.title)"
      >

        <v-card
            :color="cardColors[data.title]"
            class="pa-0 white--text"
            style="height: 100%;"
        >

          <v-card-title
              class="text-uppercase text-subtitle-2 font-weight-bold"
              v-if="data.title !== 'Image' && data.title !== 'Description'"
          >
            {{ data.title }}
          </v-card-title>

          <v-card-text class="d-flex justify-center text-h6">

            <template v-if="data.title === 'Website'">
              <a :href="data.value" target="_blank">{{ data.value }}</a>
            </template>

            <template v-else-if="data.title === 'Image'">
              <img :src="data.value" alt="Company Image" style="max-width: 20%;" />
            </template>

            <template v-else-if="data.title === 'Trading'">

              <span :class="{online: data.value, offline: !data.value}">
                {{ data.value ? 'Online' : 'Offline' }}
              </span>

            </template>

            <template v-else>
              {{ data.value }}
            </template>

          </v-card-text>

        </v-card>

      </v-col>

    </v-row>

  </v-container>

</template>

<script setup lang="ts">
import { VCard, VCardTitle, VCardText, VContainer, VRow, VCol, VTextField, VBtn } from 'vuetify/components';
import { useRoute } from 'vue-router';
import { onMounted, ref } from "vue";
import router from "@/router";

const companyData = ref([
  { title: 'Company Name', value: '' },
  { title: 'Risk Score', value: 0 },
  { title: 'Industry', value: '' },
  { title: 'Description', value: '' },
  { title: 'Price', value: 0 },
  { title: 'Changes', value: 0 },
  { title: 'Website', value: '' },
  { title: 'Image', value: '' },
  { title: 'Exchange', value: '' },
  { title: 'Is Actively Trading', value: false }
]);

const cardColors = {
  'Company Name': 'blue',
  'Financial Risk Score': 'red',
  'Industry': 'green',
  'Description': 'brown',
  'Price': 'orange',
  'Changes': 'pink',
  'Website': 'teal',
  'Image': 'grey',
  'Exchange': 'cyan',
  'Trading': 'amber'
};

const route = useRoute();

onMounted(async () => {
  const symbol = route.query.symbol as string;
  const response = await fetch(`http://localhost:8080/calculate-risk-score?symbol=${symbol}`);
  const data = await response.json();

  console.log(data);

  companyData.value = [
    { title: 'Company Name', value: data.companyName },
    { title: 'Image', value: data.image },
    { title: 'Industry', value: data.industry },
    { title: 'Description', value: data.description },
    { title: 'Price', value: `${data.price} $` },
    { title: 'Changes', value: `${data.changes} %` },
    { title: 'Financial Risk Score', value: data.riskScore },
    { title: 'Exchange', value: data.exchange },
    { title: 'Trading', value: data.isActivelyTrading },
    { title: 'Website', value: data.website },

  ];
});

const getCols = (title: string) => {
  return title === 'Description' ? 12 : 4;
};

const symbol = ref('');
function search() {
  router.push({
    path: '/calculate-risk-score',
    query: {
      symbol: symbol.value
    }
  });
  router.afterEach(() => {
    location.reload();
  });
}
</script>

<style>
.wallboard {
  padding: 20px;
  width: 90%;
  background-color:#424242;
}

.online {
  margin-right: 5px;
}

.offline {
  margin-right: 5px;
}
</style>
