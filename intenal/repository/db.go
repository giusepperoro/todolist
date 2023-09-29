package repository

import "context"

type AssociationRepo interface {
}

func (a *AssociationRepository) Create(ctx context.Context, asos *Association) error {
	return nil
}

func (a *AssociationRepository) Read(ctx context.Context, asos *Association) error {
	return nil
}

func (a *AssociationRepository) Update(ctx context.Context, asos *Association) error {
	return nil
}

func (a *AssociationRepository) Delete(ctx context.Context, asos *Association) error {
	return nil
}
