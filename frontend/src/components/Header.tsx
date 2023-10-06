import React, { HTMLAttributes, ReactPropTypes } from 'react'

interface Props extends HTMLAttributes<HTMLDivElement> {

}

export default function Header({
  className
}: Props ) {
  return (
    <h1 className={className}>
      Realtime Chat App
    </h1>
  )
}
