import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import api from '@/utils/api';

export const useCharacterStore = defineStore('character', () => {
  const characters = ref([]); // 用户所有角色
  const currentCharacter = ref(null); // 当前选中角色

  const hasCharacter = computed(() => !!currentCharacter.value);

  async function fetchCharacters() {
    const res = await api.get('/characters');
    characters.value = res || [];
    return res;
  }

  async function createCharacter(name, slotIndex) {
    const res = await api.post('/characters', { name, slot_index: slotIndex });
    await fetchCharacters();
    return res;
  }

  async function deleteCharacter(characterId) {
    await api.delete(`/characters/${characterId}`);
    await fetchCharacters();
    // 如果删除的是当前选中的角色，清除选中状态
    if (currentCharacter.value && currentCharacter.value.id === characterId) {
      currentCharacter.value = null;
      localStorage.removeItem('currentCharacterId');
    }
  }

  function selectCharacter(character) {
    currentCharacter.value = character;
    if (character) {
      localStorage.setItem('currentCharacterId', character.id);
    } else {
      localStorage.removeItem('currentCharacterId');
    }
  }

  // 根据 localStorage 恢复选中的角色
  async function restoreCharacter() {
    const savedId = localStorage.getItem('currentCharacterId');
    if (savedId) {
      await fetchCharacters();
      const found = characters.value.find((c) => c.id === parseInt(savedId));
      if (found) {
        currentCharacter.value = found;
        return found;
      }
    }
    return null;
  }

  // 获取角色详情
  async function fetchCharacterDetail(characterId) {
    const res = await api.get(`/characters/${characterId}`);
    return res;
  }

  // 清除状态
  function clearState() {
    characters.value = [];
    currentCharacter.value = null;
    localStorage.removeItem('currentCharacterId');
  }

  return {
    characters,
    currentCharacter,
    hasCharacter,
    fetchCharacters,
    createCharacter,
    deleteCharacter,
    selectCharacter,
    restoreCharacter,
    fetchCharacterDetail,
    clearState,
  };
});
