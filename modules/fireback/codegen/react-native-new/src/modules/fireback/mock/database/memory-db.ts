import {uuidv4} from '../../hooks/api';
import {BaseEntity} from '../../sdk/core/definitions';

export class MemoryEntity<T extends BaseEntity> {
  constructor(private content: T[]) {}

  items(): T[] {
    return this.content;
  }

  create(entity: Partial<T>): T {
    const newT = {
      ...entity,
      uniqueId: uuidv4().substr(0, 12),
    } as T;

    this.content.push(newT);
    return newT;
  }

  getOne(uniqueId: string): T | null {
    return this.content.find(item => item.uniqueId === uniqueId) as any;
  }

  deletes(uniqueId: string[]): boolean {
    this.content = this.content.filter(
      item => !uniqueId.includes(item.uniqueId as any),
    );

    return true;
  }

  patchOne(entity: Partial<T>): T | null {
    this.content = this.content.map(item => {
      if (item.uniqueId === entity.uniqueId) {
        return {
          ...item,
          ...entity,
        };
      }

      return item;
    });

    return entity as T;
  }
}

export const QueryToId = (name: string) => {
  return name.split(' or ').map(item => {
    return item.split(' = ')[1].trim();
  });
};
