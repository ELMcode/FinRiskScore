<template>
  <!-- Barre de recherche -->
  <v-row class="d-flex flex-wrap">
    <v-col cols="6">

      <v-text-field
          clearable
          label="Search stock symbol"
          v-model="symbol"
          @keyup.enter="search"
      />
     <!--  v-model="symbol" -> lie le champ de texte à la variable reactive 'symbol' -->
    </v-col>

    <!-- Bouton vers lien listes symbol  -->
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

  <!-- Conteneur d'informations de l'entreprise -->
  <v-container class="wallboard" fluid>
    <v-row class="d-flex flex-wrap">
      <!-- Boucle sur les données de l'entreprise pour afficher les card -->
      <v-col
          v-for="(data, index) in companyData"
          :key="data.title"
          :cols="getCols(data.title)"
      >
        <v-card
            :color="cardColors[data.title]"
            class="pa-2 white--text"
            style="height: 100%;"
        >
          <!-- titre de la card sauf pour 'Image' et 'Description' -->
          <v-card-title
              class="text-uppercase text-subtitle-2 font-weight-bold"
              v-if="data.title !== 'Image' && data.title !== 'Description'"
          >
            {{ data.title }}
          </v-card-title>

          <!-- Contenu de la card -->
          <v-card-text class="d-flex justify-center text-h6">
            <!-- Affichage en fonction du titre de la donnée -->
            <template v-if="data.title === 'Website'">
              <a :href="data.value" target="_blank">{{ data.value }}</a>
            </template>
            <template v-else-if="data.title === 'Image'">
              <img :src="data.value" alt="Company Image" style="max-width: 25%;" />
            </template>
            <template v-else-if="data.title === 'Trading'">
              <span>
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

// Données de l'entreprise
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

const route = useRoute();

// Au montage du composant
onMounted(async () => {
  const symbol = route.query.symbol as string;
  const response = await fetch(`http://localhost:8080/calculate-risk-score?symbol=${symbol}`);
  const data = await response.json();

  // Mise à jour des données de l'entreprise
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

// Fonction pour déterminer le nombre de colonnes en fonction du titre (utile pour 'Description')
const getCols = (title: string) => {
  return title === 'Description' ? 12 : 4;
};

// crée une variable réactive 'symbol', utilisé pour récupérer la valeur du texte de la barre de recherche
const symbol = ref('');
// Fonction de recherche
function search() {
  // Redirige vers  nouvelle URL avec le chemin '/calculate-risk-score' + symbole saisi de la barre de recherche via la variable reactive 'symbol'
  router.push({
    path: '/calculate-risk-score',
    query: {
      symbol: symbol.value
    }
  });
  // Rechargement de la page apres une modification de route (suite à la barre de recherche)
  router.afterEach(() => {
    location.reload();
  });
}

// Couleurs des card en fonction du titre
const cardColors = {
  'Company Name': '#5e94e3',
  'Financial Risk Score': '#D9463F',
  'Industry': '#caec6d',
  'Description': '#b8cfe0',
  'Price': '#f3c13e',
  'Changes': '#9FC131',
  'Website': '#92d0f5',
  'Image': '#E6E6E6',
  'Exchange': '#5FA1D4',
  'Trading': '#81CF95'
};
</script>

<style>
.wallboard {
  padding: 25px;
  max-width: 80%;
  background-color:#F4F8FB;
}
</style>
